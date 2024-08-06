package cmd

import (
	"fmt"
	"io"
	"log/slog"

	"github.com/charmbracelet/log"
	"github.com/gabe565/ransom/internal/ransom"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:  "ransom string...",
		Args: cobra.MinimumNArgs(1),
		RunE: run,

		SilenceErrors: true,
	}
}

func run(cmd *cobra.Command, args []string) error {
	slog.SetDefault(slog.New(log.New(cmd.ErrOrStderr())))

	result := ransom.Default().Replace(args...)
	_, _ = io.WriteString(cmd.OutOrStdout(), result+"\n")

	cmd.SilenceUsage = true
	if err := clipboard.Init(); err != nil {
		return fmt.Errorf("failed to copy to clipboard: %w", err)
	}

	clipboard.Write(clipboard.FmtText, []byte(result))
	slog.Info("Copied to clipboard")
	return nil
}
