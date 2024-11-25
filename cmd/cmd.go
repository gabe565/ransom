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
	"gabe565.com/utils/termx"
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
	conf, err := config.Load(cmd)
	if err != nil {
		return err
	}

	replacer := ransom.Default(conf.Prefix)
	var result strings.Builder
	if len(args) != 0 {
		result.WriteString(replacer.Replace(args...))
		if result.Len() != 0 {
			if _, err := io.WriteString(cmd.OutOrStdout(), result.String()+"\n"); err != nil {
				return err
			}
		}
	} else {
		if termx.IsTerminal(cmd.InOrStdin()) {
			return ErrArgs
		}

		if f, ok := cmd.InOrStdin().(*os.File); ok {
			if info, err := f.Stat(); err == nil {
				result.Grow(int(info.Size()))
			}
		}

		scanner := bufio.NewScanner(cmd.InOrStdin())
		for scanner.Scan() {
			replaced := replacer.Replace(scanner.Text() + "\n")
			if _, err := io.WriteString(cmd.OutOrStdout(), replaced); err != nil {
				return err
			}
			result.WriteString(replaced)
		}
		if scanner.Err() != nil {
			return scanner.Err()
		}
	}

	if !conf.NoCopy && result.Len() != 0 {
		cmd.SilenceUsage = true
		if err := clipboard.Init(); err != nil {
			return fmt.Errorf("failed to copy: %w", err)
		}

		clipboard.WriteText(result.String())
		slog.Info("Copied to clipboard")
	}

	return nil
}
