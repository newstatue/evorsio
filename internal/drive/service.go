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

func (s *Service) ListEntries(ctx context.Context, req ListEntriesReq) (common.PageResult[*Entry], error) {
	req.Init()
	dir := req.Dir
	size := req.Size + 1
	fsReq := &fsgen.ListEntriesRequest{
		Directory:         dir,
		Prefix:            req.Name,
		StartFromFileName: req.Cursor,
		Limit:             uint32(size),
		OmitChunks:        true,
	}

	stream, err := s.filer.ListEntries(ctx, fsReq)
	if err != nil {
		return common.PageResult[*Entry]{}, err
	}

	entries := make([]*Entry, 0, size)
	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return common.PageResult[*Entry]{}, err
		}
		entries = append(entries, NewEntryFromFS(dir, resp.GetEntry()))
	}

	hasMore := len(entries) > req.Size
	if hasMore {
		entries = entries[:req.Size]
	}

	var nextCursor string
	if len(entries) > 0 {
		nextCursor = entries[len(entries)-1].Name
	}

	return common.PageResult[*Entry]{
		Items:      entries,
		NextCursor: nextCursor,
		HasMore:    hasMore,
	}, nil
}
