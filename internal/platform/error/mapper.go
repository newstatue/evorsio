package error

import (
	"context"
	"database/sql"
	"errors"
	"evorsio/internal/auth"
	"log/slog"

	"github.com/danielgtaylor/huma/v2"
	"github.com/redis/go-redis/v9"
)

func Convert(
	ctx context.Context,
	logger *slog.Logger,
	err error,
) error {
	switch {
	case errors.Is(err, auth.ErrCodeExpired):
		return huma.Error401Unauthorized(
			"verification code expired",
		)

	case errors.Is(err, auth.ErrInvalidCode):
		return huma.Error401Unauthorized(
			"invalid verification code",
		)

	case errors.Is(err, auth.ErrUserInactive):
		return huma.Error403Forbidden(
			"user is inactive",
		)

	case errors.Is(err, sql.ErrNoRows):
		return huma.Error404NotFound(
			"resource not found",
		)

	case errors.Is(err, context.DeadlineExceeded):
		return huma.Error504GatewayTimeout(
			"request timed out",
		)

	case errors.Is(err, redis.Nil):
		return huma.Error404NotFound(
			"resource not found",
		)

	default:
		logger.ErrorContext(
			ctx,
			"unexpected application error",
			"error", err,
		)

		return huma.Error500InternalServerError(
			"unexpected error occurred",
		)
	}
}
