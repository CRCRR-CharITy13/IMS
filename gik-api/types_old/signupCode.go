package types

import "gorm.io/gorm"

type SignupCode struct {
	gorm.Model
	Code               string `json:"code"`
	Expiration         int64  `json:"expiration"`
	DesignatedUsername string `json:"designatedUsername" gorm:"unique"`
	//CreatedByUserID    uint   `json:"createdByUserID"` // do a FOREIGN KEY reference!!!
	//CreatedByUser      User   `gorm:"foreignKey:CreatedByUserID;references:ID" json:"-"`
	Expired   bool  `json:"expired"`
	CreatedAt int64 `json:"createdAt"`
}
