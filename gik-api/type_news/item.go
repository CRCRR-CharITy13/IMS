package type_news

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	SKU           string
	Name          string
	StockTotal    int
	Price         float32
	Size          string
	Warehouses    []Warehouse //`gorm:"many2many:warehouses;"`
	OrderItems    []OrderItem
	DonationItems []DonationItem
}
