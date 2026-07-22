package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func Load() (*Config, error) {
	// 本地开发时读取 .env。
	// 文件不存在时不报错，因为生产环境通常直接注入环境变量。
	_ = godotenv.Load()

	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, fmt.Errorf("parse configuration: %w", err)
	}

	return &cfg, nil
}
