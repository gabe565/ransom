package config

import (
	"gabe565.com/utils/cobrax"
	"gabe565.com/utils/must"
	"github.com/spf13/cobra"
)

const (
	FlagNoCopy = "no-copy"
	FlagPrefix = "prefix"
)

func (c *Config) RegisterFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	must.Must(cobrax.RegisterCompletionFlag(cmd))
	fs.BoolVarP(&c.NoCopy, FlagNoCopy, "n", c.NoCopy, "Disable copying to the clipboard")
	fs.StringVarP(&c.Prefix, FlagPrefix, "p", c.Prefix, `Letter prefix (can be alphabet-white or alphabet-yellow if Slack packs are enabled)`)
}
