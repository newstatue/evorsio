package main

import (
	"errors"
	"evorsio/internal/app"
	"evorsio/internal/auth"
	"evorsio/internal/platform"
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
)

func main() {
	application := bootstrap()
	defer application.DB.Close()

	rootRouter := chi.NewRouter()
	apiRouter := chi.NewRouter()

	rootRouter.Mount("/api", apiRouter)

	humaCfg := huma.DefaultConfig("Evorsio API", "1.0.0")
	humaCfg.CreateHooks = nil
	humaCfg.Servers = []*huma.Server{
		{
			URL: "/api",
		},
	}

	api := humachi.New(apiRouter, humaCfg)

	jwtService := platform.NewJWTService(
		application.Config.App.JWTSecret,
		application.Config.App.JWTIssuer,
		application.Config.App.JWTExpire,
	)
	auth.Register(
		api,
		application.Config,
		application.DB,
		application.Cache,
		application.Logger,
		jwtService,
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
			"failed to start server",
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
