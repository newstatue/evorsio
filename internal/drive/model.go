package drive

import "github.com/newstatue/evorsio/internal/resource"

type EntryType string

const (
	File   EntryType = "file"
	Folder EntryType = "folder"
)

type Entry struct {
	resource.Resource

	Type EntryType
}
