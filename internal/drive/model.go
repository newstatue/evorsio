package drive

import (
	"github.com/newstatue/evorsio/internal/fsgen"
	"github.com/newstatue/evorsio/internal/resource"
)

type EntryType string

const (
	File   EntryType = "file"
	Folder EntryType = "folder"
)

const (
	ExtendedDriveName = "evorsio.drive.name"
)

type Entry struct {
	*resource.Resource

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
	if driveName := string(entry.GetExtended()[ExtendedDriveName]); driveName != "" {
		name = driveName
	}
	return &Entry{
		Resource: resource.NewFromFS(dir, entry),
		Type:     typ,
		Name:     name,
		Mime:     entry.GetAttributes().GetMime(),
	}
}
