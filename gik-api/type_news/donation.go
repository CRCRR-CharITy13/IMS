package type_news

import (
	"gorm.io/gorm"
)

type Donation struct {
	gorm.Model
	DonorBy    uint `gorm:"foreignKey:DonorID;references:ID"`
	SignedBy   uint `gorm:"foreignKey:UserID;references:ID"`
	TotalValue float32
	TotalItems []DonationItem
}

type DonationItem struct {
	gorm.Model
	DonationId uint
	ItemID     uint
	count      uint
}
