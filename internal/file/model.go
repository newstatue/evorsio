package file

import (
	"time"

	"github.com/google/uuid"
)

type Type string

const (
	TypeFile      Type = "file"
	TypeDirectory Type = "directory"
)

type Status string

const (
	StatusPending Status = "pending"
	StatusActive  Status = "active"
	StatusDeleted Status = "deleted"
)

type File struct {
	ID       uuid.UUID
	OwnerID  uuid.UUID
	ParentID *uuid.UUID
	ObjectID *uuid.UUID

	Name   string
	Type   Type
	Status Status

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewFile(
	ownerID uuid.UUID,
	parentID *uuid.UUID,
	objectID uuid.UUID,
	name string,
) *File {
	now := time.Now().UTC()
	return &File{
		ID:       uuid.New(),
		OwnerID:  ownerID,
		ParentID: parentID,
		ObjectID: &objectID,

		Name:      name,
		Type:      TypeFile,
		Status:    StatusActive,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func NewDirectory(
	ownerID uuid.UUID,
	parentID *uuid.UUID,
	name string,
) *File {
	now := time.Now().UTC()
	return &File{
		ID:       uuid.New(),
		OwnerID:  ownerID,
		ParentID: parentID,

		Name:   name,
		Type:   TypeDirectory,
		Status: StatusActive,

		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (f *File) IsFile() bool {
	return f.Type == TypeFile
}

func (f *File) IsDirectory() bool {
	return f.Type == TypeDirectory
}

func (f *File) IsActive() bool {
	return f.Status == StatusActive
}
