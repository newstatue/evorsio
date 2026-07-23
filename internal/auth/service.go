package auth

import (
	"context"
	"crypto/rand"
	"evorsio/internal/platform/config"
	"evorsio/internal/shared"
	"evorsio/internal/user"
	"fmt"
	"log/slog"
	"math/big"
	"time"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	config   *config.Config
	logger   *slog.Logger
	cache    *redis.Client
	userRepo user.Repository
}

func NewService(config *config.Config, logger *slog.Logger, cache *redis.Client, userRepo user.Repository) *Service {
	return &Service{
		config:   config,
		logger:   logger,
		cache:    cache,
		userRepo: userRepo,
	}
}

func (s *Service) SendCode(ctx context.Context, email string) error {
	n, err := rand.Int(rand.Reader, big.NewInt(100_000))
	if err != nil {
		return fmt.Errorf("generate verification code: %w", err)
	}

	code := fmt.Sprintf("%05d", n.Int64())
	key := shared.KeyAuthCode(email)

	ttl := time.Duration(s.config.App.AuthCodeExpireSeconds) * time.Second

	if err := s.cache.Set(ctx, key, code, ttl).Err(); err != nil {
		s.logger.ErrorContext(
			ctx,
			"failed to cache verification code",
			"email", email,
			"error", err,
		)

		return fmt.Errorf("cache verification code: %w", err)
	}

	s.logger.InfoContext(
		ctx,
		"verification code cached successfully",
		"email", email,
	)

	return nil
}
