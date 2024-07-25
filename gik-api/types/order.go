package types

import "time"

type Order struct {
	OrderID          int          `gorm:"primaryKey" json:"order_id"`
	ClientID         int          `json:"client_id"`
	OrderDate        time.Time    `json:"order_date"`
	TotalCreditValue int          `json:"total_credit_value"`
	UserID           int          `json:"user_id"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
	DeletedAt        *time.Time   `json:"deleted_at"`
	Client           Client       `gorm:"foreignKey:ClientID" json:"client"`
	User             User         `gorm:"foreignKey:UserID" json:"user"`
	OrderItems       []OrdersItem `gorm:"foreignKey:OrderID" json:"order_items,omitempty"`
}
