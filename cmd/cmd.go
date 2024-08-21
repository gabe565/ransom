package cmd

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/charmbracelet/log"
	"github.com/gabe565/ransom/internal/config"
	"github.com/gabe565/ransom/internal/ransom"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ransom string...",
		RunE:    run,
		Version: buildVersion(),

		SilenceErrors: true,
	}
	cmd.SetVersionTemplate(`{{with .Name}}{{printf "%s " .}}{{end}}{{printf "commit %s" .Version}}
`)
	conf := config.New()
	conf.RegisterFlags(cmd)
	cmd.SetContext(config.NewContext(context.Background(), conf))
	return cmd
}

var ErrArgs = errors.New("requires at least one argument")

func run(cmd *cobra.Command, args []string) error {
	slog.SetDefault(slog.New(log.New(cmd.ErrOrStderr())))

	conf, ok := config.FromContext(cmd.Context())
	if !ok {
		panic("command missing config")
	}

	var result string
	if len(args) != 0 {
		result = ransom.Default().Replace(args...)
	} else {
		if f, ok := cmd.InOrStdin().(*os.File); ok {
			if isatty.IsTerminal(f.Fd()) || isatty.IsCygwinTerminal(f.Fd()) {
				return ErrArgs
			}
		}

		b, err := io.ReadAll(cmd.InOrStdin())
		if err != nil {
			return err
		}
		b = bytes.TrimSpace(b)

		result = ransom.Default().Replace(string(b))
	}

	if len(result) == 0 {
		return nil
	}

	if _, err := io.WriteString(cmd.OutOrStdout(), result+"\n"); err != nil {
		return err
	}

	if !conf.NoCopy {
		cmd.SilenceUsage = true
		if err := clipboard.Init(); err != nil {
			return fmt.Errorf("failed to copy to clipboard: %w", err)
		}

		clipboard.Write(clipboard.FmtText, []byte(result))
		slog.Info("Copied to clipboard")
	}

	return nil
}
