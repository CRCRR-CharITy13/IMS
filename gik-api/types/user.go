package types

import "time"

type User struct {
	UserID       int       `json:"user_id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	RegisteredAt time.Time `json:"registered_at"`
	IsAdmin      bool      `json:"is_admin"`
	IsDisabled   bool      `json:"is_disabled"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
