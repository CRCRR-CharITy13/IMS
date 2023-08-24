package type_news

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string
	Password    string
	RegisterdAt int64
	Admin       bool
	Disabled    bool
}
