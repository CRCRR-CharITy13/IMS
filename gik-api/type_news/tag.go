package type_news

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string `json:"name"`
}
