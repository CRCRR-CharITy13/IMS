package type_news

import (
	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model
	ItemID     uint `gorm:"foreignKey:ItemID, references:ID"`
	LocationID uint `gorm:"foreignKey:LocationID, references:ID"`
	Stock      int
}
