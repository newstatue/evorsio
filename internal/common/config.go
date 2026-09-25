package common

import (
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	App  *AppConfig  `env:",init"`
	HTTP *HTTPConfig `env:",init"`
	DB   *DBConfig   `env:",init"`
	FS   *FSConfig   `env:",init"`
}

type AppConfig struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

type HTTPConfig struct {
	Addr string `env:"ADDR" envDefault:":8080"`
}

type DBConfig struct {
	Driver string `env:"DRIVER" envDefault:"sqlite"`
	DSN    string `env:"DSN"`
}

type FSConfig struct {
	Path     string `env:"FS_PATH"`
	DataDir  string `env:"FS_DATA_DIR"`
	Addr     string `env:"FS_ADDR" envDefault:"127.0.0.1:18888"`
	MountDir string `env:"FS_MOUNT_DIR"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	if cfg.DB.DSN == "" {
		file, err := xdg.DataFile("evorsio/data.sqlite")
		if err != nil {
			return nil, err
		}
		cfg.DB.DSN = file
	}

	if cfg.FS.Path == "" {
		cfg.FS.Path = GetExePath("weed")
	}

	if cfg.FS.DataDir == "" {
		file, err := xdg.DataFile("evorsio/weed-data/.keep")
		if err != nil {
			return nil, err
		}
		cfg.FS.DataDir = filepath.Dir(file)
	}

	if cfg.FS.MountDir == "" {
		cfg.FS.MountDir = GetMountDir()
	}

	return cfg, nil
}
