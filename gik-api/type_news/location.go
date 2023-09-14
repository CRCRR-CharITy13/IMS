package type_news

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Name        string
	Description string
	Warehouses  []Warehouse //`gorm:"many2many:warehouses;"`
}
