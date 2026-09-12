package seaweedfs

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

type Server struct {
	mu sync.Mutex

	cmd  *exec.Cmd
	done chan struct{}
	l    *slog.Logger
}

func NewServer(l *slog.Logger) *Server {
	return &Server{l: l}
}

func (s *Server) Start(ctx context.Context, path string, dataDir string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cmd != nil && s.cmd.Process != nil {
		return fmt.Errorf("SeaweedFS 已经启动")
	}

	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		return fmt.Errorf("创建 SeaweedFS 数据目录失败: %w", err)
	}

	cmd := exec.CommandContext(ctx, path, "-log_json", "server", "-dir="+dataDir, "-filer", "-master.raftHashicorp")
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
	done := make(chan struct{})
	s.done = done

	go pipeLog(ctx, s.l, stdout)
	go pipeLog(ctx, s.l, stderr)
	go func() {
		defer close(done)
		err := cmd.Wait()
		s.mu.Lock()
		defer s.mu.Unlock()
		if s.cmd != cmd {
			return
		}
		s.cmd = nil
		if err != nil {
			s.l.ErrorContext(ctx, "SeaweedFS Server 已退出", KErr, err)
		}
	}()
	return nil
}

func (s *Server) Stop() error {
	s.mu.Lock()

	if s.cmd == nil || s.cmd.Process == nil {
		s.mu.Unlock()
		return nil
	}

	process, done := s.cmd.Process, s.done
	s.mu.Unlock()

	var err error
	if runtime.GOOS == "windows" {
		err = process.Kill()
	} else {
		err = process.Signal(os.Interrupt)
	}
	if err != nil && !errors.Is(err, os.ErrProcessDone) {
		return err
	}
	select {
	case <-done:
		return nil
	case <-time.After(5 * time.Second):
		if err := process.Kill(); err != nil && !errors.Is(err, os.ErrProcessDone) {
			return err
		}
		<-done
		return nil
	}
}

func (s *Server) Running() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.cmd != nil && s.cmd.Process != nil
}

func pipeLog(ctx context.Context, logger *slog.Logger, r io.Reader) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		ln := scanner.Text()

		var data map[string]any
		if err := json.Unmarshal([]byte(ln), &data); err != nil {
			logger.DebugContext(ctx, ln)
			continue
		}

		level := parseLevel(data["level"])
		msg, _ := data["msg"].(string)

		attrs := []slog.Attr{}

		if file, ok := data["file"].(string); ok {
			attrs = append(attrs, slog.String("file", file))
		}

		if line, ok := data["line"]; ok {
			attrs = append(attrs, slog.Any("line", line))
		}

		logger.LogAttrs(ctx, level, msg, attrs...)
	}

	if err := scanner.Err(); err != nil {
		logger.ErrorContext(ctx, "读取 SeaweedFS 日志失败", KErr, err)
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
