package api

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/newstatue/evorsio/internal/common"
)

type API struct {
	logger *slog.Logger
	config *common.Config
	server *http.Server
	db     *sql.DB
}

type Options struct {
	Logger *slog.Logger
	Config *common.Config
	DB     *sql.DB
}

func New(opt Options) *API {
	logger := opt.Logger
	if logger == nil {
		logger = slog.Default()
	}
	return &API{
		logger: logger,
		config: opt.Config,
		server: &http.Server{
			Addr: opt.Config.HTTP.Addr,
		},
		db: opt.DB,
	}
}

func (a *API) Run(ctx context.Context) error {
	errCh := make(chan error, 1)
	go func() {
		a.logger.InfoContext(ctx, "API 服务启动", "Addr", a.server.Addr)
		err := a.server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()
	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		a.logger.InfoContext(ctx, "API 服务关闭")
		sdCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if err := a.server.Shutdown(sdCtx); err != nil {
			return err
		}
		return nil
	}
}
