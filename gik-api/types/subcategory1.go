package types

type Subcategory1 struct {
	Subcategory1Code int    `gorm:"primaryKey" json:"subcategory1_code"`
	Subcategory1Name string `json:"subcategory1_name"`
}
