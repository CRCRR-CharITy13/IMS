package types

import "time"

type ItemsCreditValueHistory struct {
	ItemFamilyID int       `gorm:"primaryKey" json:"item_family_id"`
	Date         time.Time `json:"date"`
	CreditValue  int       `json:"credit_value"`
}
