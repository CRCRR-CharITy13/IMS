package types

import (
	"gorm.io/gorm"
)

type Donor struct {
	gorm.Model
	Name        string
	PhoneNumber string
	Address     string
	Email       string
}
