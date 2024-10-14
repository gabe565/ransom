package main

import (
	"log/slog"
	"os"

	"gabe565.com/ransom/cmd"
	"github.com/charmbracelet/log"
)

func main() {
	slog.SetDefault(slog.New(log.New(os.Stderr)))
	if err := cmd.New().Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
