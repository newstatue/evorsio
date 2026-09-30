package vault

import (
	"path"

	"github.com/newstatue/evorsio/internal/fsgen"
	"github.com/newstatue/evorsio/internal/resource"
)

type Entry struct {
	*resource.Resource
	Payload *Payload
}

type Index struct {
	*resource.Resource
	Target string
}

type Path string

const (
	PathData      Path = "/vault/data"
	PathIndexHost Path = "/vault/index/host"
	PathVault     Path = "/vault/.vault"
)

type Payload struct {
	Title    string `json:"title"`
	Username string `json:"username"`
	Password string `json:"password"`
	URL      string `json:"url"`
	Notes    string `json:"notes"`
}

func NewEntry(payload *Payload) (*Entry, error) {
	id := resource.NewID()
	r := resource.NewWithID(id, PathData.Join(id+".age"), resource.KindVault)

	return &Entry{
		Resource: r,
		Payload:  payload,
	}, nil
}

func NewHostIndexFromEntry(entry *Entry, host string) *Index {
	id := resource.NewID()

	return &Index{
		Resource: resource.NewWithID(id, PathIndexHost.Join(host, entry.ID), resource.KindVault),
		Target:   entry.Path,
	}
}

func NewEntryFromFS(entry *fsgen.Entry, payload *Payload) *Entry {
	return &Entry{
		Resource: resource.NewFromFS(string(PathData), entry),
		Payload:  payload,
	}
}

func (p Path) Join(elem ...string) string {
	parts := append([]string{string(p)}, elem...)
	return path.Join(parts...)
}
