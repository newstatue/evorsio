//go:build !dev

package common

import (
	"log/slog"

	"github.com/adrg/xdg"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(level slog.Level) {
	fn, _ := xdg.StateFile("evorsio/evorsio.log")

	writer := &lumberjack.Logger{
		Filename:   fn,
		MaxSize:    20,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}

	slog.SetDefault(slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{Level: level})))

	slog.Info("logger 初始化完成")
}
