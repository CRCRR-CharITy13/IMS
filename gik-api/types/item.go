package types

import "time"

type Item struct {
	ItemFamilyID      int        `gorm:"primaryKey" json:"item_family_id"`
	ItemSize          string     `gorm:"primaryKey" json:"item_size"`
	ItemTotalQuantity int        `json:"item_total_quantity"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	ItemFamily        ItemFamily `gorm:"foreignKey:ItemFamilyID" json:"item_family"`
}
