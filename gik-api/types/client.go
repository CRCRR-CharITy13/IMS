package types

import "time"

type Client struct {
	ClientID        int        `gorm:"primaryKey" json:"client_id"`
	Name            string     `json:"name"`
	Phone           string     `json:"phone"`
	Email           string     `json:"email"`
	Address         string     `json:"address"`
	RemainingCredit int        `json:"remaining_credit"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	Orders          []Order    `gorm:"foreignKey:ClientID" json:"orders,omitempty"`
}
