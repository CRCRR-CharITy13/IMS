package types

type ItemsDescription struct {
	DescriptionCode int    `gorm:"primaryKey" json:"description_code"`
	DescriptionName string `json:"description_name"`
}
