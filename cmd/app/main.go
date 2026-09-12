package main

import (
	"database/sql"
	"log/slog"

	"github.com/newstatue/evorsio"
	"github.com/newstatue/evorsio/internal/common"
	"github.com/newstatue/evorsio/internal/constant"
	"github.com/newstatue/evorsio/internal/seaweedfs"
	"github.com/wailsapp/wails/v3/pkg/application"
	_ "modernc.org/sqlite"
)

func init() {
	common.InitLogger(slog.LevelDebug)
}

const (
	kErr            = string(constant.LogArgError)
	kComponent      = string(constant.LogArgComponent)
	vComponentApp   = string(constant.ComponentApp)
	vComponentFS    = string(constant.ComponentFS)
	vComponentWails = string(constant.ComponentWails)
)

func main() {
	l := slog.Default()
	al := l.With(kComponent, vComponentApp)

	cfg, err := common.NewConfig()
	if err != nil {
		l.Error(string(constant.ErrParseConfig), kErr, err)
		return
	}

	db, err := sql.Open(cfg.DB.Driver, cfg.DB.DSN)
	if err != nil {
		l.Error("数据库初始化失败", kErr, err)
		return
	}
	fs := seaweedfs.New(&cfg.FS, l.With(kComponent, vComponentFS))
	defer func(db *sql.DB, fs *seaweedfs.SeaweedFS) {
		_ = db.Close()
		_ = fs.Close()
	}(db, fs)

	app := application.New(application.Options{
		Name:        "app",
		Description: "A demo of using raw HTML & CSS",
		Logger:      l.With(kComponent, vComponentWails),
		Services:    []application.Service{},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(evorsio.Assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})
	menu := app.NewMenu()

	fileMenu := menu.AddSubmenu("文件")
	fileMenu.Add("打开")
	fileMenu.Add("退出")

	app.Menu.Set(menu)

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:   "Evorsio",
		Width:   1000,
		Height:  618,
		Windows: WindowsWindow,
		Mac:     MacWindow,
		URL:     "/",
	})

	app.OnShutdown(func() {
		_ = db.Close()
		_ = fs.Close()
	})

	ctx := app.Context()

	if err := db.PingContext(ctx); err != nil {
		al.ErrorContext(ctx, "数据库连接失败", kErr, err)
		return
	}

	if err := fs.Start(ctx); err != nil {
		al.ErrorContext(ctx, "对象存储启动失败", kErr, err)
		return
	}

	if err := app.Run(); err != nil {
		al.Error("APP 退出异常", kErr, err)
		return
	}

	al.Info("APP 正常退出")
}
