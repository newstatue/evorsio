package platform

import (
	"errors"
	"fmt"
	"time"

	gojwt "github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token expired")
)

type Claims struct {
	UserID string `json:"userId"`

	gojwt.RegisteredClaims
}

type JWTService struct {
	secret     []byte
	issuer     string
	expiration time.Duration
}

func NewJWTService(
	secret string,
	issuer string,
	expiration time.Duration,
) *JWTService {
	return &JWTService{
		secret:     []byte(secret),
		issuer:     issuer,
		expiration: expiration,
	}
}

func (s *JWTService) GenerateToken(
	userID string,
) (string, error) {
	now := time.Now()

	claims := Claims{
		UserID: userID,
		RegisteredClaims: gojwt.RegisteredClaims{
			Issuer:    s.issuer,
			Subject:   userID,
			IssuedAt:  gojwt.NewNumericDate(now),
			NotBefore: gojwt.NewNumericDate(now),
			ExpiresAt: gojwt.NewNumericDate(now.Add(s.expiration)),
		},
	}

	token := gojwt.NewWithClaims(
		gojwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString(s.secret)
	if err != nil {
		return "", fmt.Errorf("sign jwt token: %w", err)
	}

	return tokenString, nil
}

func (s *JWTService) ParseToken(tokenString string) (*Claims, error) {
	token, err := gojwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *gojwt.Token) (any, error) {
			if token.Method != gojwt.SigningMethodHS256 {
				return nil, fmt.Errorf(
					"unexpected signing method: %s",
					token.Method.Alg(),
				)
			}

			return s.secret, nil
		},
		gojwt.WithIssuer(s.issuer),
		gojwt.WithExpirationRequired(),
		gojwt.WithValidMethods([]string{
			gojwt.SigningMethodHS256.Alg(),
		}),
	)

	if err != nil {
		if errors.Is(err, gojwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}

		return nil, fmt.Errorf("%w: %v", ErrInvalidToken, err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
