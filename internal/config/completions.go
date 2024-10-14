package config

import (
	"errors"

	"github.com/spf13/cobra"
)

func (c *Config) RegisterCompletions(cmd *cobra.Command) {
	if err := errors.Join(
		cmd.RegisterFlagCompletionFunc(FlagCompletion, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
			return []string{"bash", "zsh", "fish", "powershell"}, cobra.ShellCompDirectiveNoFileComp
		}),
		cmd.RegisterFlagCompletionFunc(FlagPrefix, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
			return []string{"alphabet-white", "alphabet-yellow"}, cobra.ShellCompDirectiveNoFileComp
		}),
	); err != nil {
		panic(err)
	}
}
