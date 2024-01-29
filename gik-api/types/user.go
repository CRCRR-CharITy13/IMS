package types

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username        string
	Password        string
	RegisteredAt    int64
	Admin           bool
	Disabled        bool
	SignedOrders    []Order
	SignedDonations []Donation
}
