package config

import "github.com/spf13/cobra"

const (
	NoCopyFlag = "no-copy"
)

func (c *Config) RegisterFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVarP(&c.NoCopy, NoCopyFlag, "n", c.NoCopy, "Disable copying to the clipboard")
}
