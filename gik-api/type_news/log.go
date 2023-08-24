package type_news

import (
	"gorm.io/gorm"
)

type AdvancedLog struct {
	gorm.Model
	IPAddress string `json:"ipAddress"`
	UserAgent string `json:"userAgent"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	UserID    uint   `json:"userId"`
	User      User   `gorm:"foreignKey:UserID;references:ID" json:"-"`
	Timestamp int64  `json:"timestamp"`
}

type SimpleLog struct {
	gorm.Model
	UserID    uint   `json:"userId"`
	User      User   `gorm:"foreignKey:UserID;references:ID" json:"-"`
	Action    string `json:"action"`
	Timestamp int64  `json:"timestamp"`
	IPAddress string `json:"ipAddress"`
}
