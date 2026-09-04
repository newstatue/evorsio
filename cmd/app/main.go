package main

import (
	"database/sql"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
	"github.com/newstatue/evorsio"
	"github.com/newstatue/evorsio/internal/common"
	"github.com/wailsapp/wails/v3/pkg/application"
	_ "modernc.org/sqlite"
)

func main() {
	l := slog.New(tint.NewTextHandler(colorable.NewColorable(os.Stderr), &tint.Options{
		Level:      slog.LevelInfo,
		TimeFormat: time.RFC3339Nano,
	}))
	slog.SetDefault(l)
	cfg, err := common.NewConfig()
	if err != nil {
		l.Error("配置出错", "error", err)
		return
	}

	db, err := sql.Open("sqlite", cfg.DB.DSN)
	if err != nil {
		l.Error("数据库初始化失败", "error", err)
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	app := application.New(application.Options{
		Name:        "app",
		Description: "A demo of using raw HTML & CSS",
		Logger:      l,
		Services:    []application.Service{},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(evorsio.Assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  "Window 1",
		Width:  1000,
		Height: 618,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(6, 7, 15),
		URL:              "/",
	})

	if err := db.PingContext(app.Context()); err != nil {
		l.ErrorContext(app.Context(), "数据库连接失败", "error", err)
		return
	}

	if err := app.Run(); err != nil {
		l.Error("APP 退出异常", "error", err)
		return
	}

	l.Info("APP 正常退出")
}
