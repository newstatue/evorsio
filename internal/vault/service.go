package vault

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log/slog"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/newstatue/evorsio/internal/common"
	"github.com/newstatue/evorsio/internal/fsgen"
	"github.com/newstatue/evorsio/internal/resource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	l         *slog.Logger
	vault     *Vault
	filer     fsgen.SeaweedFilerClient
	generator *Generator
}

func NewService(l *slog.Logger, vault *Vault, filer fsgen.SeaweedFilerClient, generator *Generator) *Service {
	return &Service{
		vault:     vault,
		filer:     filer,
		generator: generator,
		l:         l,
	}
}

func (s *Service) IsInitialized(ctx context.Context) (bool, error) {
	_, err := s.filer.LookupDirectoryEntry(
		ctx,
		&fsgen.LookupDirectoryEntryRequest{
			Directory: path.Dir(string(PathVault)),
			Name:      path.Base(string(PathVault)),
		},
	)
	if err == nil {
		return true, nil
	}

	s.l.ErrorContext(ctx, "报错", "err", err)
	if status.Code(err) == codes.Unknown {
		return false, nil
	}

	return false, err
}

func (s *Service) IsLocked() bool {
	return s.vault.IsLocked()
}

func (s *Service) Init(ctx context.Context, masterPass string) error {
	initialized, err := s.IsInitialized(ctx)
	if err != nil {
		return err
	}
	if initialized {
		return nil
	}

	data, err := s.vault.Init(masterPass)
	if err != nil {
		return err
	}

	now := time.Now()
	_, err = s.filer.CreateEntry(
		ctx,
		&fsgen.CreateEntryRequest{
			Directory: path.Dir(string(PathVault)),
			Entry: &fsgen.Entry{
				Name:    path.Base(string(PathVault)),
				Content: data,
				Attributes: &fsgen.FuseAttributes{
					FileMode: 0644,
					FileSize: uint64(len(data)),
					Mtime:    now.Unix(),
					Crtime:   now.Unix(),
				},
			},
		})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Unlock(ctx context.Context, masterPass string) error {
	resp, err := s.filer.LookupDirectoryEntry(
		ctx,
		&fsgen.LookupDirectoryEntryRequest{
			Directory: path.Dir(string(PathVault)),
			Name:      path.Base(string(PathVault)),
		},
	)
	if err != nil {
		return err
	}

	entry := resp.GetEntry()

	return s.vault.Unlock(masterPass, entry.GetContent())
}

type GeneratePasswordReq struct {
	Length    int
	Uppercase bool
	Lowercase bool
	Numbers   bool
	Symbols   bool
}

func (s *Service) GeneratePassword(ctx context.Context, req GeneratePasswordReq) (string, error) {

	return s.generator.Generate(ctx, Options{
		Length:    req.Length,
		Uppercase: req.Uppercase,
		Lowercase: req.Lowercase,
		Numbers:   req.Numbers,
		Symbols:   req.Symbols,
	})
}

type CreateEntryReq struct {
	Title    string
	Username string
	Password string
	URL      string
	Notes    string
}

func (s *Service) CreateEntry(ctx context.Context, req CreateEntryReq) error {
	entry, err := NewEntry(&Payload{
		Title:    req.Title,
		Username: req.Username,
		Password: req.Password,
		URL:      req.URL,
		Notes:    req.Notes,
	})
	if err != nil {
		return err
	}

	err = s.createEntry(ctx, entry)
	if err != nil {
		return err
	}

	err = s.createIndex(ctx, entry)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) createEntry(ctx context.Context, entry *Entry) error {
	data, err := s.vault.Encrypt(entry.Payload)
	if err != nil {
		return err
	}
	now := time.Now()

	filerReq := &fsgen.CreateEntryRequest{
		Directory: path.Dir(entry.Path),
		Entry: &fsgen.Entry{
			Name:    path.Base(entry.Path),
			Content: data,
			Attributes: &fsgen.FuseAttributes{
				FileMode: 0644,
				FileSize: uint64(len(data)),
				Mtime:    now.Unix(),
				Crtime:   now.Unix(),
			},
			Extended: map[string][]byte{
				resource.ExtendedID:   []byte(entry.ID),
				resource.ExtendedKind: []byte(entry.Kind),
			},
		},
	}
	_, err = s.filer.CreateEntry(ctx, filerReq)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) createIndex(ctx context.Context, entry *Entry) error {
	u, err := url.Parse(entry.Payload.URL)
	if err != nil {
		return err
	}
	now := time.Now()

	host := strings.ToLower(u.Hostname())
	if host != "" {
		index := NewHostIndexFromEntry(entry, host)
		_, err = s.filer.CreateEntry(ctx, &fsgen.CreateEntryRequest{
			Directory: path.Dir(index.Path),
			Entry: &fsgen.Entry{
				Name: path.Base(index.Path),
				Attributes: &fsgen.FuseAttributes{
					FileMode:      0644,
					SymlinkTarget: index.Target,
					Mtime:         now.Unix(),
					Crtime:        now.Unix(),
				},
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

type ListEntriesReq struct {
	common.PageQuery
	Host string
}

func (s *Service) ListEntries(ctx context.Context, req ListEntriesReq) (common.PageResult[*Entry], error) {
	req.Init()

	size := req.Size + 1
	fsReq := &fsgen.ListEntriesRequest{
		Directory:         string(PathData),
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
		entry := resp.GetEntry()

		payload, err := s.vault.Decrypt(bytes.NewReader(entry.GetContent()))
		if err != nil {
			return common.PageResult[*Entry]{}, err
		}

		entries = append(entries, NewEntryFromFS(entry, payload))
	}

	hasMore := len(entries) > req.Size
	if hasMore {
		entries = entries[:req.Size]
	}

	var nextCursor string
	if len(entries) > 0 {
		nextCursor = path.Base(entries[len(entries)-1].Path)
	}

	return common.PageResult[*Entry]{
		Items:      entries,
		NextCursor: nextCursor,
		HasMore:    hasMore,
	}, nil
}
