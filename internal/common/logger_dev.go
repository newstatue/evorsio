//go:build dev

package common

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
)

func InitLogger() error {
	l := slog.New(tint.NewTextHandler(colorable.NewColorable(os.Stderr), &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.RFC3339,
	}))
	slog.SetDefault(l)

	return nil
}
