package resource

import (
	"path"
	"time"
	"uuid"

	"github.com/newstatue/evorsio/internal/fsgen"
)

type Kind string

const (
	ExtendedID   = "evorsio.id"
	ExtendedKind = "evorsio.kind"
)

const (
	KindDrive Kind = "drive"
	KindVault Kind = "vault"
)

type Resource struct {
	ID        string
	Kind      Kind
	Path      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(dir string, kind Kind) *Resource {

	id := uuid.NewV7().String()
	return &Resource{
		ID:   id,
		Kind: kind,
		Path: path.Join(string(kind), dir, id),
	}
}

func NewFromFS(dir string, entry *fsgen.Entry) *Resource {
	attr := entry.GetAttributes()

	return &Resource{
		ID:        string(entry.GetExtended()[ExtendedID]),
		Kind:      Kind(entry.GetExtended()[ExtendedKind]),
		Path:      path.Join(dir, entry.GetName()),
		CreatedAt: time.Unix(attr.GetCrtime(), int64(attr.GetCrtimeNs())),
		UpdatedAt: time.Unix(attr.GetMtime(), int64(attr.GetMtimeNs())),
	}
}
