package types

import "time"

type Donor struct {
	DonorID   int        `gorm:"primaryKey" json:"donor_id"`
	Name      string     `json:"name"`
	Phone     string     `json:"phone"`
	Email     string     `json:"email"`
	Address   string     `json:"address"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Donations []Donation `gorm:"foreignKey:DonorID" json:"donations,omitempty"`
}
