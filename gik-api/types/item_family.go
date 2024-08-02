package types

import "time"

type ItemFamily struct {
	ItemFamilyID       int       `gorm:"primaryKey" json:"item_family_id"`
	CategoryCode       string    `json:"category_code"`
	Subcategory1Code   int       `json:"subcategory1_code"`
	Subcategory2Code   string    `json:"subcategory2_code"`
	DescriptionCode    int       `json:"description_code"`
	ItemName           string    `json:"item_name"`
	ItemSKU            string    `json:"item_sku"`
	CurrentCreditValue int       `json:"current_credit_value"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}