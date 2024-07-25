package types

import "time"

type OrdersItem struct {
	OrderID      int        `gorm:"primaryKey" json:"order_id"`
	ItemFamilyID int        `gorm:"primaryKey" json:"item_family_id"`
	ItemSize     string     `gorm:"primaryKey" json:"item_size"`
	Quantity     int        `json:"quantity"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Order        Order      `gorm:"foreignKey:OrderID" json:"order"`
	ItemFamily   ItemFamily `gorm:"foreignKey:ItemFamilyID" json:"item_family"`
	Size         Size       `gorm:"foreignKey:ItemSize" json:"size"`
}
