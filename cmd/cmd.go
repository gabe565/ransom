package cmd

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"gabe565.com/ransom/internal/clipboard"
	"gabe565.com/ransom/internal/config"
	"gabe565.com/ransom/internal/ransom"
	"gabe565.com/utils/cobrax"
	"github.com/charmbracelet/log"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

func New(opts ...cobrax.Option) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "ransom string...",
		RunE: run,

		ValidArgsFunction: cobra.NoFileCompletions,
		SilenceErrors:     true,
	}
	conf := config.New()
	conf.RegisterFlags(cmd)
	conf.RegisterCompletions(cmd)
	cmd.SetContext(config.NewContext(context.Background(), conf))
	for _, opt := range opts {
		opt(cmd)
	}
	return cmd
}

var ErrArgs = errors.New("requires at least one argument")

func run(cmd *cobra.Command, args []string) error {
	slog.SetDefault(slog.New(log.New(cmd.ErrOrStderr())))

	conf, ok := config.FromContext(cmd.Context())
	if !ok {
		panic("command missing config")
	}

	if conf.Prefix != "" {
		conf.Prefix += "-"
	}
	replacer := ransom.Default(conf.Prefix)
	var result string
	if len(args) != 0 {
		result = replacer.Replace(args...)
		if len(result) != 0 {
			if _, err := io.WriteString(cmd.OutOrStdout(), result+"\n"); err != nil {
				return err
			}
		}
	} else {
		if f, ok := cmd.InOrStdin().(*os.File); ok && isatty.IsTerminal(f.Fd()) || isatty.IsCygwinTerminal(f.Fd()) {
			return ErrArgs
		}

		scanner := bufio.NewScanner(cmd.InOrStdin())
		for scanner.Scan() {
			replaced := replacer.Replace(scanner.Text() + "\n")
			if _, err := io.WriteString(cmd.OutOrStdout(), replaced); err != nil {
				return err
			}
			result += replaced
		}
		if scanner.Err() != nil {
			return scanner.Err()
		}
		result = strings.TrimRight(result, "\n")
	}

	if len(result) != 0 && !conf.NoCopy {
		cmd.SilenceUsage = true
		if err := clipboard.Init(); err != nil {
			return fmt.Errorf("failed to copy: %w", err)
		}

		clipboard.WriteText(result)
		slog.Info("Copied to clipboard")
	}

	return nil
}
