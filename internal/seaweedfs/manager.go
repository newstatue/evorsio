package seaweedfs

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os/exec"

	"github.com/newstatue/evorsio/internal/common"
	"github.com/newstatue/evorsio/internal/constant"
	"github.com/newstatue/evorsio/internal/fsgen"
)

const (
	KErr                = string(constant.LogArgError)
	KSubComponent       = string(constant.LogArgSubComponent)
	VSubComponentServer = string(constant.SubComponentServer)
	VSubComponentMount  = string(constant.SubComponentMount)
)

type Manager struct {
	l   *slog.Logger
	cfg *common.FSConfig

	mount  *Server
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
		mount:  NewServer(l.With(KSubComponent, VSubComponentMount)),
		client: client,
	}, nil
}

func (m *Manager) Start(ctx context.Context) error {
	cmdS := exec.CommandContext(
		ctx,
		m.cfg.Path,
		"-log_json",
		"server",
		"-dir="+m.cfg.DataDir,
		"-filer",
		"-master.raftHashicorp",
		"-webdav",
	)
	configureProcess(cmdS)

	if err := m.server.Start(ctx, cmdS); err != nil {
		return fmt.Errorf("启动 SeaweedFS server: %w", err)
	}

	if err := m.client.WaitReady(ctx); err != nil {
		_ = m.server.Close()
		return fmt.Errorf("等待 SeaweedFS filer 就绪: %w", err)
	}

	cmdM := exec.CommandContext(
		ctx,
		m.cfg.Path,
		"-log_json",
		"mount",
		"-filer=127.0.0.1:8888",
		"-dir="+m.cfg.MountDir,
		"-volumeName=evorsio",
	)
	configureProcess(cmdM)

	if err := m.mount.Start(ctx, cmdM); err != nil {
		_ = m.server.Close()
		return fmt.Errorf("挂载 SeaweedFS: %w", err)
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

	if err := m.mount.Close(); err != nil {
		errs = append(errs, fmt.Errorf("关闭 SeaweedFS mount: %w", err))
	}
	return errors.Join(errs...)
}

func (m *Manager) Filer() fsgen.SeaweedFilerClient {
	return m.client.filer
}
