package auth

import "errors"

var (
	ErrorInvalidCode  = errors.New("invalid verification code")
	ErrorCodeExpired  = errors.New("verification code expired")
	ErrorUserInactive = errors.New("user is inactive")
)
