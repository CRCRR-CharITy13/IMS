package type_news

import (
	"gorm.io/gorm"
)

type Donation struct {
	gorm.Model
	DonorID    uint
	DonorBy    Donor `gorm:"foreignKey:DonorID;references:ID"`
	UserID     uint
	SignedBy   User `gorm:"foreignKey:UserID;references:ID"`
	TotalValue float32
	TotalItems []DonationItem
}

type DonationItem struct {
	gorm.Model
	DonationID uint
	Donation   Donation `gorm:"foreignKey:DonationID;references:ID"`
	ItemID     uint
	Item       Item `gorm:"foreignKey:ItemID;references:ID"`
	count      uint
}
