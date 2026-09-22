package resource

import (
	"time"
	"uuid"
)

type Kind string

const (
	KindDrive Kind = "drive"
)

type Resource struct {
	ID        string
	Kind      Kind
	Path      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(path string, kind Kind) (Resource, error) {
	now := time.Now()

	return Resource{
		ID:        uuid.NewV7().String(),
		Kind:      kind,
		Path:      path,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
