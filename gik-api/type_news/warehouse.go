package type_news

import (
	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model
	ItemID     uint
	Item       Item `gorm:"foreignKey:ItemID;references:ID"`
	LocationID uint
	Location   Location `gorm:"foreignKey:LocationID;references:ID"`
	Stock      int
}
