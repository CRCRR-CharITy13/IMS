package types

import "time"

type Location struct {
	LocationID     int             `gorm:"primaryKey" json:"location_id"`
	Type           string          `json:"type"`
	Site           string          `json:"site"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeletedAt      *time.Time      `json:"deleted_at"`
	ItemsLocations []ItemsLocation `gorm:"foreignKey:LocationID" json:"items_locations,omitempty"`
}
