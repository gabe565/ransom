package cmd

import (
	"io"

	"github.com/gabe565/ransom/internal/ransom"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:  "ransom string...",
		Args: cobra.MinimumNArgs(1),
		RunE: run,
	}
}

func run(cmd *cobra.Command, args []string) error {
	_, _ = io.WriteString(cmd.OutOrStdout(), ransom.Default().Replace(args...)+"\n")
	return nil
}
