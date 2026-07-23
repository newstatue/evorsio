package user

import (
	"context"
	"database/sql"
	"evorsio/internal/shared"

	"github.com/google/uuid"
)

type Repository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*shared.User, error)
	FindByEmail(ctx context.Context, email string) (*shared.User, error)
	Create(ctx context.Context, user *shared.User) (*shared.User, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) FindByID(ctx context.Context, id uuid.UUID) (*shared.User, error) {
	const query = `
	select id, email, coalesce(name,""), status, created_at, updated_at
	from users
	where id = $1
	`

	user := &shared.User{}

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresRepository) FindByEmail(ctx context.Context, email string) (*shared.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresRepository) Create(ctx context.Context, user *shared.User) (*shared.User, error) {
	//TODO implement me
	panic("implement me")
}
