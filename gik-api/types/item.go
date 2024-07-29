package types

import "time"

type Item struct {
	ItemFamilyID      int       `json:"item_family_id"`
	ItemSize          string    `json:"item_size"`
	ItemTotalQuantity int       `json:"item_total_quantity"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
