package types

import "time"

type AdvancedLog struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	IPAddress string     `json:"ip_address"`
	UserAgent string     `json:"user_agent"`
	Method    string     `json:"method"`
	Path      string     `json:"path"`
	Timestamp int        `json:"timestamp"`
	UserID    int        `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	User      User       `gorm:"foreignKey:UserID" json:"user"`
}
