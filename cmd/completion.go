package cmd

import "github.com/spf13/cobra"

func generateCompletion(cmd *cobra.Command, shell string) error {
	switch shell {
	case "bash":
		return cmd.GenBashCompletion(cmd.OutOrStdout())
	case "zsh":
		return cmd.GenZshCompletion(cmd.OutOrStdout())
	case "fish":
		return cmd.GenFishCompletion(cmd.OutOrStdout(), true)
	case "powershell":
		return cmd.GenPowerShellCompletion(cmd.OutOrStdout())
	default:
		return nil
	}
}
