package types

import "time"

type User struct {
	UserID       int        `gorm:"primaryKey" json:"user_id"`
	Username     string     `json:"username"`
	Password     string     `json:"password"`
	RegisteredAt int        `json:"registered_at"`
	IsAdmin      bool       `json:"is_admin"`
	IsDisabled   bool       `json:"is_disabled"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Sessions     []Session  `gorm:"foreignKey:UserID" json:"sessions,omitempty"`
	Orders       []Order    `gorm:"foreignKey:UserID" json:"orders,omitempty"`
	Donations    []Donation `gorm:"foreignKey:UserID" json:"donations,omitempty"`
}
