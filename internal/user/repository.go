package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Repository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) (*User, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) FindByID(ctx context.Context, id uuid.UUID) (*User, error) {
	const query = `
	select id, email, coalesce(name,''), status, created_at, updated_at
	from users
	where id = $1
	`

	user := &User{}

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to find user by id: %w", err)
	}

	return user, nil
}

func (r *PostgresRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	const query = `
	select id, email, coalesce(name,''), status, created_at, updated_at
	from users
	where email = $1
	`
	user := &User{}

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}

	return user, nil
}

func (r *PostgresRepository) Create(ctx context.Context, user *User) (*User, error) {
	const query = `
	insert into users (id, email, name, status, created_at, updated_at)
	values ($1, $2, $3, $4, $5, $6)
	returning id, email, coalesce(name,''), status, created_at, updated_at
    `

	newUser := &User{}

	err := r.db.QueryRowContext(ctx, query,
		user.ID, user.Email, user.Name,
		user.Status, user.CreatedAt, user.UpdatedAt,
	).Scan(
		&newUser.ID,
		&newUser.Email,
		&newUser.Name,
		&newUser.Status,
		&newUser.CreatedAt,
		&newUser.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	return newUser, nil
}
