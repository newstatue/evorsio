package shared

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/go-chi/jwtauth/v5"
)

const APIAuthScheme = "Auth"

func APIAuthSchemes() map[string]*huma.SecurityScheme {
	return map[string]*huma.SecurityScheme{
		APIAuthScheme: {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}
}

func APIRequireAuth(api huma.API) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		token, claims, err := jwtauth.FromContext(ctx.Context())
		if err != nil || token == nil {
			_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, "invalid or expired token")
			return
		}

		userID, ok := claims["sub"].(string)
		if !ok || userID == "" {
			_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, "invalid token")
			return
		}

		next(ctx)
	}
}
