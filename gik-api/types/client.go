package types

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name    string  `json:"name"`
	Contact string  `json:"contact"`
	Phone   string  `json:"phone"`
	Email   string  `json:"email"`
	Address string  `json:"address"`
	Balance float32 `json:"balance"`
}
