package types

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name     string  `json:"name"`
	SKU      string  `json:"sku"`
	Category string  `json:"category"` //Gender and such
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
	Size     string  `json:"size"`
}
