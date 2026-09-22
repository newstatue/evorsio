package drive

import (
	"database/sql"

	"github.com/newstatue/evorsio/internal/dbgen"
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
