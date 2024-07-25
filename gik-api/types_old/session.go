package types

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	ID        string `json:"id"`
	UserID    uint   `json:"userId"`
	User      User   `json:"user" gorm:"foreignKey:UserID;references:ID"`
	CreatedAt int64  `json:"createdAt"`
	ExpiresAt int64  `json:"expiresAt"`
}
