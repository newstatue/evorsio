package main

import (
	"context"
	"database/sql"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/newstatue/evorsio/internal/api"
	"github.com/newstatue/evorsio/internal/common"
	_ "modernc.org/sqlite"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	l := slog.New(slog.NewTextHandler(os.Stdout, nil))
	cfg, err := common.NewConfig()
	if err != nil {
		l.ErrorContext(ctx, "配置出错", "error", err)
		os.Exit(1)
	}

	db, err := sql.Open("sqlite", cfg.DB.DSN)
	if err != nil {
		l.ErrorContext(ctx, "数据库初始化失败", "error", err)
		os.Exit(1)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	if err := db.PingContext(ctx); err != nil {
		l.ErrorContext(ctx, "数据库连接失败", "error", err)
		os.Exit(1)
	}
	a := api.New(api.Options{
		Logger: l,
		Config: cfg,
		DB:     db,
	})

	if err := a.Run(ctx); err != nil {
		l.Error("API 退出异常", "error", err)
		os.Exit(1)
	}

	l.Info("应用正常退出")
}
