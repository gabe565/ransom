package config

import (
	"github.com/spf13/cobra"
)

const (
	FlagCompletion = "completion"
	FlagNoCopy     = "no-copy"
	FlagPrefix     = "prefix"
)

func (c *Config) RegisterFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.StringVar(&c.Completion, FlagCompletion, c.Completion, "Generate the autocompletion script for the specified shell (one of bash, zsh, fish, powershell)")
	fs.BoolVarP(&c.NoCopy, FlagNoCopy, "n", c.NoCopy, "Disable copying to the clipboard")
	fs.StringVarP(&c.Prefix, FlagPrefix, "p", c.Prefix, `Letter prefix (can be alphabet-white or alphabet-yellow if Slack packs are enabled)`)
}
