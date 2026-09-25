package seaweedfs

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"os/exec"
	"sync"
	"syscall"
)

type Server struct {
	mu  sync.Mutex
	cmd *exec.Cmd
	l   *slog.Logger
}

var (
	ErrProcessAlreadyRunning = errors.New("进程已经在运行")
)

func NewServer(l *slog.Logger) *Server {
	return &Server{l: l}
}

func (s *Server) Start(ctx context.Context, c *exec.Cmd) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cmd != nil && s.cmd.Process != nil {
		return ErrProcessAlreadyRunning
	}

	// 启动新进程
	cmd := c
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}

	s.cmd = cmd

	go pipeLog(s.l, stdout)
	go pipeLog(s.l, stderr)

	go func() {
		err := cmd.Wait()
		if s.cmd != cmd {
			return
		}
		s.cmd = nil
		if err != nil && ctx.Err() == nil {
			s.l.ErrorContext(ctx, "Manager Server 已退出", KErr, err)
		}
	}()
	return nil
}

func pipeLog(logger *slog.Logger, r io.Reader) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		ln := scanner.Text()

		var data map[string]any
		if err := json.Unmarshal([]byte(ln), &data); err != nil {
			logger.Debug(ln)
			continue
		}

		level := parseLevel(data["level"])
		msg, _ := data["msg"].(string)

		var attrs []slog.Attr

		if file, ok := data["file"].(string); ok {
			attrs = append(attrs, slog.String("file", file))
		}

		if line, ok := data["line"]; ok {
			attrs = append(attrs, slog.Any("line", line))
		}

		logger.LogAttrs(context.Background(), level, msg, attrs...)
	}

	if err := scanner.Err(); err != nil {
		logger.Error("读取 Manager 日志失败", KErr, err)
	}
}

func parseLevel(v any) slog.Level {
	level, _ := v.(string)

	switch level {
	case "DEBUG":
		return slog.LevelDebug
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func (s *Server) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cmd == nil || s.cmd.Process == nil {
		return nil
	}

	err := s.cmd.Process.Signal(syscall.SIGTERM)
	if err != nil {
		err = s.cmd.Process.Kill()
	}
	s.cmd = nil
	return err
}
