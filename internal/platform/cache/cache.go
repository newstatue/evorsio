package cache

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

func New(uri string) (*redis.Client, error) {
	if strings.TrimSpace(uri) == "" {
		return nil, errors.New("cache empty uri")
	}
	opt, err := redis.ParseURL(uri)
	if err != nil {
		return nil, err
	}

	opt.DialTimeout = 5 * time.Second
	opt.ReadTimeout = 3 * time.Second
	opt.WriteTimeout = 3 * time.Second

	rdb := redis.NewClient(opt)

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
