package object

import (
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, object *Object) error
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository { return &PostgresRepository{db: db} }

func (r *PostgresRepository) Create(ctx context.Context, object *Object) error {
	return nil
}
