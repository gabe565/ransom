package config

import (
	"io"
	"log/slog"
	"time"

	"gabe565.com/utils/termx"
	"github.com/lmittmann/tint"
)

func InitLog(w io.Writer) {
	slog.SetDefault(slog.New(tint.NewHandler(w, &tint.Options{
		Level:      slog.LevelInfo,
		TimeFormat: time.Kitchen,
		NoColor:    !termx.IsColor(w),
	})))
}
