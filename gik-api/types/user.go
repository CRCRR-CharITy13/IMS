package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username        string `json:"username" gorm:"unique"`
	Password        string `json:"-"` // hashed with argon2id
	TwoFactorSecret string `json:"-"`
	RegisteredAt    int64  `json:"registeredAt"`
	Admin           bool   `json:"-"`
	Disabled        bool   `json:"disabled" gorm:"default:false"`
}
