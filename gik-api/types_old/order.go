package types

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ClientID   uint
	Client     Client `gorm:"foreignKey:ClientID;references:ID"`
	UserID     uint
	SignedBy   User `gorm:"foreignKey:UserID;references:ID"`
	TotalCost  float32
	TotalItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID uint
	Order   Order `gorm:"foreignKey:OrderID;references:ID"`
	ItemID  uint
	Item    Item `gorm:"foreignKey:ItemID;references:ID"`
	Count   int
}
