package types

import "time"

type ItemsLocation struct {
	ItemFamilyID int        `gorm:"primaryKey" json:"item_family_id"`
	ItemSize     string     `gorm:"primaryKey" json:"item_size"`
	LocationID   int        `gorm:"primaryKey" json:"location_id"`
	Quantity     int        `json:"quantity"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	ItemFamily   ItemFamily `gorm:"foreignKey:ItemFamilyID" json:"item_family"`
	Size         Size       `gorm:"foreignKey:ItemSize" json:"size"`
	Location     Location   `gorm:"foreignKey:LocationID" json:"location"`
}
