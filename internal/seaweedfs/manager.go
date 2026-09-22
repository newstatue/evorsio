package seaweedfs

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/newstatue/evorsio/internal/common"
	"github.com/newstatue/evorsio/internal/constant"
)

const (
	KErr                = string(constant.LogArgError)
	KSubComponent       = string(constant.LogArgSubComponent)
	VSubComponentServer = string(constant.SubComponentServer)
)

type Manager struct {
	l   *slog.Logger
	cfg *common.FSConfig

	server *Server
	client *Client
}

func NewManager(cfg *common.FSConfig, l *slog.Logger) (*Manager, error) {
	client, err := NewClient(cfg.Addr)
	if err != nil {
		return nil, fmt.Errorf("创建 gRPC client: %w", err)
	}
	return &Manager{
		l:      l,
		cfg:    cfg,
		server: NewServer(l.With(KSubComponent, VSubComponentServer)),
		client: client,
	}, nil
}

func (m *Manager) Start(ctx context.Context) error {
	if err := m.server.Start(ctx, m.cfg.Path, m.cfg.DataDir); err != nil {
		return fmt.Errorf("启动 SeaweedFS server: %w", err)
	}
	if err := m.client.WaitReady(ctx); err != nil {
		return fmt.Errorf("等待 SeaweedFS filer 就绪: %w", err)
	}
	return nil
}

func (m *Manager) Close() error {
	var errs []error
	if err := m.client.Close(); err != nil {
		errs = append(errs, fmt.Errorf("关闭 gRPC client: %w", err))
	}
	if err := m.server.Close(); err != nil {
		errs = append(errs, fmt.Errorf("关闭 SeaweedFS server: %w", err))
	}
	return nil
}
