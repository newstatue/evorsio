package seaweedfs

import (
	"context"
	"fmt"
	"time"

	"github.com/newstatue/evorsio/internal/fsgen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultTickerTime = time.Millisecond * 200
)

type Client struct {
	conn  *grpc.ClientConn
	filer fsgen.SeaweedFilerClient
}

func NewClient(addr string) (*Client, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Client{
		conn:  conn,
		filer: fsgen.NewSeaweedFilerClient(conn),
	}, nil
}

func (c *Client) Close() error {
	if c.conn == nil {
		return nil
	}
	return c.conn.Close()
}

func (c *Client) Ping(ctx context.Context) error {
	_, err := c.filer.Ping(ctx, &fsgen.PingRequest{})
	if err != nil {
		return fmt.Errorf("filer 连接失败: %w", err)
	}
	return nil
}

func (c *Client) WaitReady(ctx context.Context) error {
	ticker := time.NewTicker(defaultTickerTime)
	defer ticker.Stop()
	for {
		if err := c.Ping(ctx); err == nil {
			return nil
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("等待 filer 就绪失败 %w", ctx.Err())
		case <-ticker.C:
		}
	}
}
