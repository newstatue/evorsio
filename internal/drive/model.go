package drive

import "github.com/newstatue/evorsio/internal/resource"

type Drive struct {
	resource.Resource

	OwnerID string
	Quota   int64
}
