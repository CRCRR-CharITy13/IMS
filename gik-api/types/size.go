package types

type Size struct {
	ItemSize    string `gorm:"primaryKey" json:"item_size"`
	Description string `json:"description"`
}
