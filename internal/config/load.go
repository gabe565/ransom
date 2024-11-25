package config

import (
	"log/slog"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

func Load(cmd *cobra.Command) (*Config, error) {
	slog.SetDefault(slog.New(log.New(cmd.ErrOrStderr())))

	conf, ok := FromContext(cmd.Context())
	if !ok {
		panic("command missing config")
	}

	if conf.Prefix != "" {
		conf.Prefix += "-"
	}
	return conf, nil
}
