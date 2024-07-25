package types

type Subcategory2 struct {
	Subcategory2Code string `gorm:"primaryKey" json:"subcategory2_code"`
	Subcategory2Name string `json:"subcategory2_name"`
}
