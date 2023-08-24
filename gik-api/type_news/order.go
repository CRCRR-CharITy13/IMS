package type_news

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ClientID   uint `gorm:"foreignKey:ClientID;references:ID"`
	SignedBy   uint `gorm:"foreignKey:UserID;references:ID"`
	TotalCost  float32
	TotalItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID uint
	ItemID  uint
	count   int
}
