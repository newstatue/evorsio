package drive

import (
	"path"

	"github.com/newstatue/evorsio/internal/fsgen"
	"github.com/newstatue/evorsio/internal/resource"
)

type EntryType string

const (
	File   EntryType = "file"
	Folder EntryType = "folder"
)

type Entry struct {
	resource.Resource

	Name string
	Type EntryType
	Mime string
}

func NewEntryFromFS(dir string, entry *fsgen.Entry) *Entry {
	var typ EntryType
	if entry.GetIsDirectory() {
		typ = Folder
	} else {
		typ = File
	}
	name := entry.GetName()
	return &Entry{
		Resource: resource.New(path.Join(dir, name), resource.KindDrive),
		Type:     typ,
		Name:     name,
		Mime:     entry.GetAttributes().GetMime(),
	}
}
