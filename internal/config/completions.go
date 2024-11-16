package config

import (
	"gabe565.com/utils/must"
	"github.com/spf13/cobra"
)

func (c *Config) RegisterCompletions(cmd *cobra.Command) {
	must.Must(cmd.RegisterFlagCompletionFunc(FlagPrefix, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{"alphabet-white", "alphabet-yellow"}, cobra.ShellCompDirectiveNoFileComp
	}))
}
