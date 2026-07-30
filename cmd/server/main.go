package main

import (
	"errors"
	"evorsio/internal/app"
	"evorsio/internal/auth"
	"evorsio/internal/platform/cache"
	"evorsio/internal/platform/config"
	"evorsio/internal/platform/database"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

func main() {
	application := bootstrap()
	defer application.DB.Close()

	tokenAuth := jwtauth.New(
		"HS256",
		[]byte(application.Config.JWT.JWTSecret),
		nil,
	)

	rootRouter := chi.NewRouter()
	apiRouter := chi.NewRouter()

	rootRouter.Use(middleware.RequestID)
	rootRouter.Use(middleware.Recoverer)

	// 只负责解析 JWT，不会强制所有接口登录。
	apiRouter.Use(jwtauth.Verifier(tokenAuth))

	rootRouter.Mount("/api", apiRouter)

	humaCfg := huma.DefaultConfig(
		"Evorsio API",
		"1.0.0",
	)
	humaCfg.CreateHooks = nil
	humaCfg.Servers = []*huma.Server{
		{
			URL: "/api",
		},
	}

	humaCfg.Components.SecuritySchemes =
		map[string]*huma.SecurityScheme{
			"bearerAuth": {
				Type:         "http",
				Scheme:       "bearer",
				BearerFormat: "JWT",
			},
		}

	api := humachi.New(apiRouter, humaCfg)

	auth.Register(
		api,
		application.Config,
		application.DB,
		application.Cache,
		application.Logger,
		tokenAuth,
	)

	addr := fmt.Sprintf(
		"%s:%d",
		application.Config.Server.Host,
		application.Config.Server.Port,
	)

	server := &http.Server{
		Addr:    addr,
		Handler: rootRouter,
	}

	application.Logger.Info(
		"starting server",
		"addr", addr,
	)

	if err := server.ListenAndServe(); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		application.Logger.Error(
			"server stopped unexpectedly",
			"error", err,
		)
	}
}

func bootstrap() *app.App {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(cfg.Database.DSN)
	if err != nil {
		log.Fatal(err)
	}

	rdb, err := cache.New(cfg.Cache.URI)
	if err != nil {
		log.Fatal(err)
	}

	logger := slog.Default()

	return app.New(cfg, db, rdb, logger)
}
