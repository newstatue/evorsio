package object

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusActive   Status = "active"
	StatusDeleting Status = "deleting"
	StatusDeleted  Status = "deleted"
)

type Object struct {
	ID uuid.UUID

	Bucket string
	Key    string

	SHA256      string
	Size        int64
	ContentType string
	ETag        string

	StorageClass string
	RefCount     int64
	Status       Status

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (o *Object) IsActive() bool {
	return o.Status == StatusActive
}

func (o *Object) HasReferences() bool {
	return o.RefCount > 0
}
