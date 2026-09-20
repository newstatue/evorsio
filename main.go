package main

import (
	"database/sql"
	"log/slog"

	"github.com/newstatue/evorsio/db"
	"github.com/newstatue/evorsio/internal/common"
	"github.com/newstatue/evorsio/internal/constant"
	"github.com/newstatue/evorsio/internal/option"
	"github.com/newstatue/evorsio/internal/seaweedfs"
	"github.com/pressly/goose/v3"
	"github.com/wailsapp/wails/v3/pkg/application"
	_ "modernc.org/sqlite"
)

const (
	kErr            = string(constant.LogArgError)
	kComponent      = string(constant.LogArgComponent)
	vComponentApp   = string(constant.ComponentApp)
	vComponentFS    = string(constant.ComponentFS)
	vComponentWails = string(constant.ComponentWails)
)

func init() {
	common.InitLogger(slog.LevelDebug)
	common.InitMigration(db.Migrations)
}

func main() {
	l := slog.Default()
	al := l.With(kComponent, vComponentApp)

	al.Info("APP 开始启动")

	cfg, err := common.NewConfig()
	if err != nil {
		l.Error(string(constant.ErrParseConfig), kErr, err)
	}

	d, err := sql.Open(cfg.DB.Driver, cfg.DB.DSN)
	if err != nil {
		l.Error("数据库初始化失败", kErr, err)
	}
	if err := goose.Up(d, "migrations"); err != nil {
		l.Error("数据库迁移失败", kErr, err)
	}
	fs := seaweedfs.New(&cfg.FS, l.With(kComponent, vComponentFS), SeaweedFS)
	defer func(db *sql.DB, fs *seaweedfs.SeaweedFS) {
		_ = db.Close()
		_ = fs.Close()
	}(d, fs)

	app := application.New(application.Options{
		Name:        "app",
		Description: "A demo of using raw HTML & CSS",
		Logger:      l.With(kComponent, vComponentWails),
		Services:    []application.Service{},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(Assets),
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

	app.Window.NewWithOptions(option.WindowOpt)

	app.OnShutdown(func() {
		_ = d.Close()
		_ = fs.Close()
	})

	ctx := app.Context()

	if err := d.PingContext(ctx); err != nil {
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
