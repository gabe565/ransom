package config

import "github.com/spf13/cobra"

func Load(cmd *cobra.Command) (*Config, error) {
	InitLog(cmd.ErrOrStderr())

	conf, ok := FromContext(cmd.Context())
	if !ok {
		panic("command missing config")
	}

	if conf.Prefix != "" {
		conf.Prefix += "-"
	}
	return conf, nil
}
