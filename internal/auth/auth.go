package auth

import (
	"database/sql"
	"evorsio/internal/platform"
	"evorsio/internal/platform/config"
	"evorsio/internal/user"
	"log/slog"

	"github.com/danielgtaylor/huma/v2"
	"github.com/redis/go-redis/v9"
)

func Register(
	api huma.API,
	config *config.Config,
	db *sql.DB,
	cache *redis.Client,
	logger *slog.Logger,
	jwt *platform.JWTService,
) {
	userRepo := user.NewPostgresRepository(db)
	authService := NewService(config, logger, cache, userRepo, jwt)
	authHandler := NewHandler(authService)

	huma.Post(api, "/auth/send-code", authHandler.SendCode)
	huma.Post(api, "/auth/login", authHandler.Login)

}
