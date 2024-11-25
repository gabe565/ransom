package main

import (
	"log/slog"
	"os"

	"gabe565.com/ransom/cmd"
	"gabe565.com/ransom/internal/config"
	"gabe565.com/utils/cobrax"
)

func main() {
	config.InitLog(os.Stderr)
	root := cmd.New(cobrax.WithVersion(""))
	if err := root.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
