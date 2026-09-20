package common

import (
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	HTTP HTTPConfig
	DB   DBConfig
	FS   FSConfig
}

type HTTPConfig struct {
	Addr string `env:"ADDR" envDefault:":8080"`
}

type DBConfig struct {
	Driver string `env:"DRIVER" envDefault:"sqlite"`
	DSN    string `env:"DSN"`
}

type FSConfig struct {
	Path    string `env:"FS_PATH"`
	DataDir string `env:"FS_DATA_DIR"`
	Addr    string `env:"FS_ADDR" envDefault:"127.0.0.1:18888"`
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

	if cfg.FS.DataDir == "" {
		cfg.FS.DataDir = filepath.Join(xdg.DataHome, "evorsio", "weed-data")
		if err != nil {
			return nil, err
		}
		cfg.FS.DataDir = file
	}

	return cfg, nil
}
