package auth

import (
	"context"
	"database/sql"
	"evorsio/internal/platform/config"
	"evorsio/internal/shared"
	"evorsio/internal/user"
	"log/slog"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/go-chi/jwtauth/v5"
	"github.com/redis/go-redis/v9"
)

func Register(
	api huma.API,
	cfg *config.Config,
	db *sql.DB,
	cache *redis.Client,
	logger *slog.Logger,
	tokenAuth *jwtauth.JWTAuth,
) {
	type TestInput struct{}

	type TestOutput struct {
		Body struct {
			UserID  string `json:"userId"`
			Message string `json:"message"`
		}
	}

	testHandler := func(ctx context.Context, input *TestInput) (*TestOutput, error) {
		_, claims, err := jwtauth.FromContext(ctx)
		if err != nil {
			return nil, huma.Error401Unauthorized(
				"unauthorized",
			)
		}

		userID, ok := claims["sub"].(string)
		if !ok || userID == "" {
			return nil, huma.Error401Unauthorized(
				"invalid user identity",
			)
		}

		output := &TestOutput{}
		output.Body.UserID = userID
		output.Body.Message = "authenticated"

		return output, nil
	}

	userRepo := user.NewPostgresRepository(db)
	authService := NewService(cfg, logger, cache, userRepo, tokenAuth)
	authHandler := NewHandler(authService)

	huma.Post(api, "/auth/send-code", authHandler.SendCode)
	huma.Post(api, "/auth/login", authHandler.Login)

	huma.Register(api, huma.Operation{
		Path:   "/auth/test",
		Method: http.MethodGet,
		Middlewares: huma.Middlewares{
			shared.APIRequireAuth(api),
		},
		Security: []map[string][]string{
			{shared.APIAuthScheme: {}},
		},
	}, testHandler)
}
