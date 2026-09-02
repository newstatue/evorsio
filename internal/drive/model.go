package drive

import "github.com/newstatue/evorsio/internal/resource"

type File struct {
	resource.Resource

	Size     int64
	MimeType string
}

type Folder struct {
	resource.Resource
}

type Symlink struct {
	resource.Resource

	TargetID string
}

type Entry struct {
	ParentID string
	ChildID  string
}

func NewFile(name string, size int64, mimeType string) File {
	r := resource.New(name, resource.TypeFile)
	return File{
		Resource: r,
		Size:     size,
		MimeType: mimeType,
	}
}

func NewFolder(name string) Folder {
	r := resource.New(name, resource.TypeFolder)
	return Folder{
		Resource: r,
	}
}

func NewSymlink(name string, targetID string) Symlink {
	r := resource.New(name, resource.TypeSymlink)
	return Symlink{
		Resource: r,
		TargetID: targetID,
	}
}
