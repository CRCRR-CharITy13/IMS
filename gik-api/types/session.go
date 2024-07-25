package types

type Session struct {
	ID        string `gorm:"primaryKey" json:"id"`
	UserID    int    `json:"user_id"`
	CreatedAt int    `json:"created_at"`
	ExpiresAt int    `json:"expires_at"`
	User      User   `gorm:"foreignKey:UserID" json:"user"`
}
