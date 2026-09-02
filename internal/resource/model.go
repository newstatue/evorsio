package resource

import (
	"time"
	"uuid"
)

type Type string

const (
	TypeFile    Type = "file"
	TypeFolder  Type = "folder"
	TypeSymlink Type = "symlink"
)

type Resource struct {
	ID        string
	Type      Type
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(name string, typ Type) Resource {
	return Resource{
		ID:        uuid.NewV7().String(),
		Name:      name,
		Type:      typ,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
}
