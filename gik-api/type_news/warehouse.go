package type_news

import (
	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model
	ItemID     uint
	LocationID uint
	Stock      int
}
