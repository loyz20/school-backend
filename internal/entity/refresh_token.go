package entity

import "time"

type RefreshToken struct {
	ID         string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	PenggunaID string `gorm:"not null"`
	Token      string `gorm:"not null"`
	IPAddress  string
	UserAgent  string
	ExpiredAt  time.Time `gorm:"not null"`
	IsRevoked  bool      `gorm:"default:false"`
	CreatedAt  time.Time
}
