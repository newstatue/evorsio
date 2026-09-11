package drive

import (
	"github.com/newstatue/evorsio/internal/resource"
)

type File struct {
	resource.Resource

	Size     int64
	MimeType string
}

func NewFile(name string, size int64, mimeType string) File {
	r := resource.New(name, resource.TypeFile)
	return File{
		Resource: r,
		Size:     size,
		MimeType: mimeType,
	}
}
