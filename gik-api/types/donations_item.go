package types

import "time"

type DonationsItem struct {
	DonationID   int        `gorm:"primaryKey" json:"donation_id"`
	ItemFamilyID int        `gorm:"primaryKey" json:"item_family_id"`
	ItemSize     string     `gorm:"primaryKey" json:"item_size"`
	Quantity     int        `json:"quantity"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Donation     Donation   `gorm:"foreignKey:DonationID" json:"donation"`
	ItemFamily   ItemFamily `gorm:"foreignKey:ItemFamilyID" json:"item_family"`
	Size         Size       `gorm:"foreignKey:ItemSize" json:"size"`
}
