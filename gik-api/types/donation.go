package types

import "time"

type Donation struct {
	DonationID        int             `gorm:"primaryKey" json:"donation_id"`
	DonorID           int             `json:"donor_id"`
	DonationDate      time.Time       `json:"donation_date"`
	TotalDollarsValue float64         `json:"total_dollars_value"`
	TotalCreditValue  int             `json:"total_credit_value"`
	UserID            int             `json:"user_id"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	DeletedAt         *time.Time      `json:"deleted_at"`
	Donor             Donor           `gorm:"foreignKey:DonorID" json:"donor"`
	User              User            `gorm:"foreignKey:UserID" json:"user"`
	DonationItems     []DonationsItem `gorm:"foreignKey:DonationID" json:"donation_items,omitempty"`
}
