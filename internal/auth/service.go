package auth

import (
	"context"
	"crypto/rand"
	"database/sql"
	"errors"
	"evorsio/internal/platform"
	"evorsio/internal/platform/config"
	"evorsio/internal/shared"
	"evorsio/internal/user"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	config     *config.Config
	logger     *slog.Logger
	cache      *redis.Client
	userRepo   user.Repository
	jwtService *platform.JWTService
}

func NewService(
	config *config.Config,
	logger *slog.Logger,
	cache *redis.Client,
	userRepo user.Repository,
	jwtService *platform.JWTService,
) *Service {
	return &Service{
		config:     config,
		logger:     logger,
		cache:      cache,
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (s *Service) SendCode(ctx context.Context, email string) error {
	n, err := rand.Int(rand.Reader, big.NewInt(100_000))
	if err != nil {
		return fmt.Errorf("generate verification code: %w", err)
	}

	code := fmt.Sprintf("%05d", n.Int64())
	key := shared.KeyAuthCode(email)

	ttl := s.config.App.AuthCodeExpire

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

func (s *Service) LoginAndReturnToken(ctx context.Context, email string, code string) (string, error) {
	cachedCode, err := s.cache.Get(ctx, shared.KeyAuthCode(email)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrorCodeExpired
		}

		s.logger.ErrorContext(
			ctx,
			"failed to get cached verification code",
			"email", email,
			"error", err,
		)

		return "", err
	}

	if cachedCode != code {
		return "", ErrorInvalidCode
	}

	_, _ = s.cache.Del(ctx, shared.KeyAuthCode(email)).Result()

	userEntity, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			userEntity, err = s.userRepo.Create(ctx, shared.NewUser(email))
			if err != nil {

				s.logger.ErrorContext(
					ctx,
					"failed to create new user",
					"email", email,
					"error", err,
				)

				return "", err
			}
		} else {
			s.logger.ErrorContext(
				ctx,
				"failed to get user by email",
				"email", email,
				"error", err,
			)
			return "", err
		}
	}

	if userEntity.Status == shared.UserStatusInactive {
		return "", ErrorUserInactive
	}

	token, err := s.jwtService.GenerateToken(userEntity.ID.String())
	if err != nil {
		s.logger.ErrorContext(
			ctx,
			"failed to get generated token",
			"userId", userEntity.ID,
			"email", email,
			"error", err,
		)
		return "", err
	}

	return token, err
}
