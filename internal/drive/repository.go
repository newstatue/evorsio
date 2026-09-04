package drive

import (
	"context"
	"database/sql"
	"uuid"

	"github.com/newstatue/evorsio/internal/dbgen"
	"github.com/newstatue/evorsio/internal/resource"
)

type Repository struct {
	q  *dbgen.Queries
	db *sql.DB
}

func NewRepository(q *dbgen.Queries) *Repository {
	return &Repository{
		q: q,
	}
}

func (r *Repository) GetFile(ctx context.Context, id uuid.UUID) (*File, error) {
	row, err := r.q.SelectFileById(ctx, id.String())
	if err != nil {
		return nil, err
	}
	return &File{
		Resource: resource.Resource{
			ID:        row.Resource.ID,
			Type:      resource.Type(row.Resource.Type),
			Name:      row.Resource.Name,
			CreatedAt: row.Resource.CreatedAt,
			UpdatedAt: row.Resource.UpdatedAt,
		},
		Size:     row.File.Size,
		MimeType: row.File.MimeType,
	}, nil
}

func (r *Repository) GetFolder(ctx context.Context, id uuid.UUID) (*Folder, error) {
	row, err := r.q.SelectFolderById(ctx, id.String())
	if err != nil {
		return nil, err
	}
	return &Folder{
		Resource: resource.Resource{
			ID:        row.Resource.ID,
			Type:      resource.Type(row.Resource.Type),
			Name:      row.Resource.Name,
			CreatedAt: row.Resource.CreatedAt,
			UpdatedAt: row.Resource.UpdatedAt,
		},
	}, nil
}

func (r *Repository) GetSymlink(ctx context.Context, id uuid.UUID) (*Symlink, error) {
	row, err := r.q.SelectSymlinkById(ctx, id.String())
	if err != nil {
		return nil, err
	}
	return &Symlink{
		Resource: resource.Resource{
			ID:        row.Resource.ID,
			Type:      resource.Type(row.Resource.Type),
			Name:      row.Resource.Name,
			CreatedAt: row.Resource.CreatedAt,
			UpdatedAt: row.Resource.UpdatedAt,
		},
		TargetID: row.Symlink.TargetID,
	}, nil
}

func (r *Repository) CreateFile(ctx context.Context, file File) (*File, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	qtx := r.q.WithTx(tx)

	rParams := dbgen.InsertResourceParams{
		ID:        file.Resource.ID,
		Type:      string(file.Resource.Type),
		Name:      file.Resource.Name,
		CreatedAt: file.Resource.CreatedAt,
		UpdatedAt: file.Resource.UpdatedAt,
	}
	if err := qtx.InsertResource(ctx, rParams); err != nil {
		return nil, err
	}

	fParams := dbgen.InsertFileParams{
		ResourceID: file.Resource.ID,
		Size:       file.Size,
		MimeType:   file.MimeType,
	}
	if err := qtx.InsertFile(ctx, fParams); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &file, nil
}

func (r *Repository) CreateFolder(ctx context.Context, folder Folder) (*Folder, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	qtx := r.q.WithTx(tx)

	rParams := dbgen.InsertResourceParams{
		ID:        folder.Resource.ID,
		Type:      string(folder.Resource.Type),
		Name:      folder.Resource.Name,
		CreatedAt: folder.Resource.CreatedAt,
		UpdatedAt: folder.Resource.UpdatedAt,
	}
	if err := qtx.InsertResource(ctx, rParams); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &folder, nil
}

func (r *Repository) CreateSymlink(ctx context.Context, symlink Symlink) (*Symlink, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	qtx := r.q.WithTx(tx)

	rParams := dbgen.InsertResourceParams{
		ID:        symlink.Resource.ID,
		Type:      string(symlink.Resource.Type),
		Name:      symlink.Resource.Name,
		CreatedAt: symlink.Resource.CreatedAt,
		UpdatedAt: symlink.Resource.UpdatedAt,
	}
	if err := qtx.InsertResource(ctx, rParams); err != nil {
		return nil, err
	}

	sParams := dbgen.InsertSymlinkParams{
		ResourceID: symlink.Resource.ID,
		TargetID:   symlink.TargetID,
	}
	if err := qtx.InsertSymlink(ctx, sParams); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &symlink, nil
}
