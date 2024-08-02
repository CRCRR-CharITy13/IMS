package types

import "time"

type SimpleLog struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	Action    string     `json:"action"`
	Timestamp int        `json:"timestamp"`
	IPAddress string     `json:"ip_address"`
	UserID    int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
