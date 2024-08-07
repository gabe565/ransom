package cmd

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/charmbracelet/log"
	"github.com/gabe565/ransom/internal/config"
	"github.com/gabe565/ransom/internal/ransom"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ransom string...",
		Args:    cobra.MinimumNArgs(1),
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

func run(cmd *cobra.Command, args []string) error {
	slog.SetDefault(slog.New(log.New(cmd.ErrOrStderr())))

	conf, ok := config.FromContext(cmd.Context())
	if !ok {
		panic("command missing config")
	}

	result := ransom.Default().Replace(args...)
	_, _ = io.WriteString(cmd.OutOrStdout(), result+"\n")

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
