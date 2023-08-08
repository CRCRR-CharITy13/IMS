package types

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Timestamp     int64  `json:"timestamp"`
	ClientID      uint   `json:"clientId"`
	Client        Client `json:"-" gorm:"foreignKey:ClientID;references:ID"`
	Type          bool   `json:"type"` //True = import False = Export
	SignerID      uint   `json:"signerId"`
	Signer        User   `json:"-" gorm:"foreignKey:SignerID;references:ID"`
	TotalQuantity int    `json:"totalQuantity"`
}

type TransactionItem struct {
	gorm.Model
	TransactionID uint        `json:"transactionId"`
	Transaction   Transaction `gorm:"foreignKey:TransactionID;references:ID" json"-"`
	Quantity      int         `json:"quantity"`
	ProductID     int         `json:"productId"`
}
