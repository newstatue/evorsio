package drive

import "github.com/newstatue/evorsio/internal/resource"

type Folder struct {
	resource.Resource
}

func NewFolder(name string) Folder {
	r := resource.New(name, resource.TypeFolder)
	return Folder{
		Resource: r,
	}
}
