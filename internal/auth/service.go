package auth

import (
	"context"
	"crypto/rand"
	"database/sql"
	"errors"
	"evorsio/internal/platform/config"
	"evorsio/internal/shared"
	"evorsio/internal/user"
	"fmt"
	"log/slog"
	"math/big"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	config    *config.Config
	logger    *slog.Logger
	cache     *redis.Client
	userRepo  user.Repository
	tokenAuth *jwtauth.JWTAuth
}

func NewService(
	config *config.Config,
	logger *slog.Logger,
	cache *redis.Client,
	userRepo user.Repository,
	tokenAuth *jwtauth.JWTAuth,
) *Service {
	return &Service{
		config:    config,
		logger:    logger,
		cache:     cache,
		userRepo:  userRepo,
		tokenAuth: tokenAuth,
	}
}

func (s *Service) SendCode(ctx context.Context, email string) error {
	n, err := rand.Int(rand.Reader, big.NewInt(100_000))
	if err != nil {
		return fmt.Errorf("generate verification code: %w", err)
	}

	code := fmt.Sprintf("%05d", n.Int64())
	key := shared.KeyAuthCode(email)

	ttl := s.config.APP.AuthCodeExpire

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
			return "", ErrInvalidCode
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
		return "", ErrInvalidCode
	}

	userEntity, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			s.logger.ErrorContext(
				ctx,
				"failed to get user by email",
				"email", email,
				"error", err,
			)
			return "", fmt.Errorf("failed to get user by email: %w", err)
		}

		userEntity = user.NewUser(email)
		err = s.userRepo.Create(ctx, userEntity)
		if err != nil {
			s.logger.ErrorContext(
				ctx,
				"failed to create new user",
				"email", email,
				"error", err,
			)
			return "", fmt.Errorf("failed to create new user: %w", err)
		}
	}

	if userEntity.Status == user.StatusInactive {
		return "", ErrUserInactive
	}

	now := time.Now()

	_, token, err := s.tokenAuth.Encode(map[string]any{
		"sub": userEntity.ID.String(),
		"iss": s.config.JWT.JWTIssuer,
		"iat": now.Unix(),
		"nbf": now.Unix(),
		"exp": now.Add(s.config.JWT.JWTExpire).Unix(),
	})
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

	err = s.cache.Del(ctx, shared.KeyAuthCode(email)).Err()
	if err != nil {
		s.logger.ErrorContext(
			ctx,
			"failed to delete verification code",
			"email", email,
			"error", err,
		)
	}

	return token, nil
}
