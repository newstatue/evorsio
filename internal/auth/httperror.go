package auth

import (
	"errors"

	"github.com/danielgtaylor/huma/v2"
)

var (
	ErrInvalidCode  = errors.New("invalid verification code")
	ErrCodeExpired  = errors.New("verification code expired")
	ErrUserInactive = errors.New("user is inactive")
)

func toHTTPError(err error) error {
	switch {
	case errors.Is(err, ErrCodeExpired):
		return huma.Error401Unauthorized("verification code expired")

	case errors.Is(err, ErrInvalidCode):
		return huma.Error401Unauthorized("invalid verification code")

	case errors.Is(err, ErrUserInactive):
		return huma.Error403Forbidden("user is inactive")

	default:
		return err
	}
}
