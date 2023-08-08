package types

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Name   string `gorm:"name" json:"name"`
	Letter string `gorm:"type:char(1)" json:"letter"`
	SKU    string `gorm:"sku" json:"sku"`
}
