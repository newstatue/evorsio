package object

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
)

var (
	ErrObjectKeyRequired  = errors.New("object key required")
	ErrObjectBodyRequired = errors.New("object body required")
)

type Storage interface {
	Upload(ctx context.Context, key string, body io.Reader, contentType string) (string, error)

	Download(ctx context.Context, key string) (io.ReadCloser, error)

	Delete(ctx context.Context, key string) error
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) Upload(ctx context.Context, key string, body io.Reader, contentType string) (string, error) {
	key = strings.TrimSpace(key)
	if key == "" {
		return "", ErrObjectKeyRequired
	}

	if body == nil {
		return "", ErrObjectBodyRequired
	}

	contentType = strings.TrimSpace(contentType)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	etag, err := s.storage.Upload(
		ctx,
		key,
		body,
		contentType,
	)
	if err != nil {
		return "", fmt.Errorf("upload object %q: %w", key, err)
	}

	return etag, nil
}

func (s *Service) Download(ctx context.Context, key string) (io.ReadCloser, error) {
	key = strings.TrimSpace(key)
	if key == "" {
		return nil, ErrObjectKeyRequired
	}

	body, err := s.storage.Download(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("download object %q: %w", key, err)
	}

	return body, nil
}

func (s *Service) Delete(ctx context.Context, key string) error {
	key = strings.TrimSpace(key)
	if key == "" {
		return ErrObjectKeyRequired
	}

	if err := s.storage.Delete(ctx, key); err != nil {
		return fmt.Errorf("delete object %q: %w", key, err)
	}

	return nil
}
