package entity

import (
	"time"
)

type RefreshToken struct {
	ID         string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	PenggunaID string `gorm:"not null"`
	Token      string `gorm:"not null"`
	IPAddress  string
	UserAgent  string
	ExpiredAt  time.Time `gorm:"not null"`
	IsRevoked  bool      `gorm:"default:false"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoCreateTime" json:"updated_at"`
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}
