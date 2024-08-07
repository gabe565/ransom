package main

import (
	"log/slog"
	"os"

	"github.com/charmbracelet/log"
	"github.com/gabe565/ransom/cmd"
)

func main() {
	slog.SetDefault(slog.New(log.New(os.Stderr)))
	if err := cmd.New().Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
