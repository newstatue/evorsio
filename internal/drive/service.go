package drive

import (
	"context"
	"errors"
	"io"

	"github.com/newstatue/evorsio/internal/common"
	"github.com/newstatue/evorsio/internal/fsgen"
)

type Service struct {
	r     *Repository
	filer fsgen.SeaweedFilerClient
}

func NewService(r *Repository, filer fsgen.SeaweedFilerClient) *Service {
	return &Service{
		r:     r,
		filer: filer,
	}
}

type ListEntriesReq struct {
	common.PageQuery
	Dir  string
	Name string
}

func (s *Service) ListEntries(ctx context.Context, req ListEntriesReq) ([]*Entry, error) {
	req.Init()
	dir := req.Dir
	fsReq := &fsgen.ListEntriesRequest{
		Directory:         dir,
		Prefix:            req.Name,
		StartFromFileName: req.Cursor,
		Limit:             uint32(req.Size),
		OmitChunks:        true,
	}

	stream, err := s.filer.ListEntries(ctx, fsReq)
	if err != nil {
		return nil, err
	}

	entries := make([]*Entry, 0, req.Size)
	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		entries = append(entries, NewEntryFromFS(dir, resp.GetEntry()))
	}

	return entries, nil
}
