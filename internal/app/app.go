package app

import (
	"database/sql"
	"evorsio/internal/platform/config"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type App struct {
	Config *config.Config
	DB     *sql.DB
	Cache  *redis.Client
	Logger *slog.Logger
}

func New(cfg *config.Config, db *sql.DB, cache *redis.Client, logger *slog.Logger) *App {
	return &App{
		Config: cfg,
		DB:     db,
		Cache:  cache,
		Logger: logger,
	}
}
