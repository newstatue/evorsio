package vault

import (
	"context"
	"path"

	"github.com/newstatue/evorsio/internal/fsgen"
	"github.com/newstatue/evorsio/internal/resource"
)

type Service struct {
	vault     *Vault
	filer     fsgen.SeaweedFilerClient
	generator *Generator
}

func NewService(vault *Vault, filer fsgen.SeaweedFilerClient, generator *Generator) *Service {
	return &Service{
		vault:     vault,
		filer:     filer,
		generator: generator,
	}
}

type CreateVaultReq struct {
	Title    string
	Username string
	Password string
	URL      string
	Notes    string
}

func (s *Service) CreateVault(ctx context.Context, req CreateVaultReq) error {
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

	data, err := s.vault.Encrypt(&entry.Payload)
	if err != nil {
		return err
	}

	filerReq := &fsgen.CreateEntryRequest{
		Directory: path.Dir(entry.Path),
		Entry: &fsgen.Entry{
			Name:    path.Base(entry.Path),
			Content: data,
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
