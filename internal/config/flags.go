package config

import "github.com/spf13/cobra"

const (
	NoCopyFlag = "no-copy"
)

func (c *Config) RegisterFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(&c.NoCopy, NoCopyFlag, c.NoCopy, "Disable copying to the clipboard")
}
