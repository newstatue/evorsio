package shared

import (
	"time"

	"github.com/google/uuid"
)

type UserStatus int

const (
	UserStatusActive UserStatus = iota
	UserStatusInactive
)

type User struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	Email     string     `db:"email" json:"email"`
	Name      string     `db:"name" json:"name"`
	Status    UserStatus `db:"status" json:"status"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
}
