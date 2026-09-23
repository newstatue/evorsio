package drive

import (
	"database/sql"

	"github.com/newstatue/evorsio/internal/dbgen"
)

type Repository struct {
	q  *dbgen.Queries
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		q:  dbgen.New(db),
		db: db,
	}
}
