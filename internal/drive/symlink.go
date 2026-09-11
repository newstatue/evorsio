package drive

import (
	"github.com/newstatue/evorsio/internal/resource"
)

type Symlink struct {
	resource.Resource

	TargetID string
}

func NewSymlink(name string, targetID string) Symlink {
	r := resource.New(name, resource.TypeSymlink)
	return Symlink{
		Resource: r,
		TargetID: targetID,
	}
}
