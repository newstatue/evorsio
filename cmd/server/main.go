package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"evorsio/internal/config"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

type Health struct {
	Message string `json:"message"`
}

type HealthResponse struct {
	Body Health
}

func main() {
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	rootRouter := chi.NewRouter()
	// 创建 Router
	apiRouter := chi.NewRouter()

	humaCfg := huma.DefaultConfig("Evorsio API", "1.0.0")
	humaCfg.CreateHooks = nil

	// 创建 Huma API
	api := humachi.New(
		apiRouter,
		humaCfg,
	)

	// 注册接口
	huma.Get(api, "/health", func(ctx context.Context, input *struct{}) (*HealthResponse, error) {
		return &HealthResponse{
			Body: Health{
				Message: "ok",
			},
		}, nil
	})

	rootRouter.Mount("/api", apiRouter)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: rootRouter,
	}

	log.Printf("Server listening on %s", addr)

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
