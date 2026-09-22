package drive

import (
	"context"

	"github.com/newstatue/evorsio/internal/common"
	"github.com/newstatue/evorsio/internal/fsgen"
)

type Service struct {
	r     *Repository
	filer fsgen.SeaweedFilerClient
}

func NewService(r *Repository) *Service {
	return &Service{
		r: r,
	}
}

func (s *Service) ListEntries(ctx context.Context, path string, request common.PageQuery) ([]Entry, error) {
	fsgen.ListEntriesRequest{
		Directory:          "",
		Prefix:             "",
		StartFromFileName:  "",
		InclusiveStartFrom: false,
		Limit:              0,
		SnapshotTsNs:       0,
		OmitChunks:         false,
	}
	s.filer.ListEntries(ctx)
}
