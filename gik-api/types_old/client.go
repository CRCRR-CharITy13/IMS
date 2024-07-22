package types

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	OrgName     string
	Address     string
	PhoneNumber string
	Email       string
	Contact     string
	Balance     float32
}
