//go:build !dev

package common

import (
	"log/slog"
	"os"

	"github.com/adrg/xdg"
	"gopkg.in/natefinch/lumberjack.v2"
)

var level slog.LevelVar

func InitLogger() {
	switch os.Getenv("LOG_LEVEL") {
	case "debug":
		level.Set(slog.LevelDebug)
	case "warn":
		level.Set(slog.LevelWarn)
	case "error":
		level.Set(slog.LevelError)
	default:
		level.Set(slog.LevelInfo)
	}

	fn, _ := xdg.StateFile("evorsio/evorsio.log")

	writer := &lumberjack.Logger{
		Filename:   fn,
		MaxSize:    20,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}

	slog.SetDefault(slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{Level: &level})))

	slog.Info("logger 初始化完成")
}
