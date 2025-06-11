package entity

import "time"

type Pengguna struct {
	PenggunaID     string  `gorm:"type:uuid;primaryKey" json:"pengguna_id"`
	SekolahID      string  `gorm:"type:uuid" json:"sekolah_id"`
	Username       string  `json:"username"`
	Nama           string  `json:"nama"`
	PeranIDStr     string  `json:"peran_id_str"`
	Password       string  `json:"-"` // tidak diserialisasi
	Alamat         string  `json:"alamat"`
	NoTelepon      string  `json:"no_telepon"`
	NoHP           string  `json:"no_hp"`
	PtkID          *string `gorm:"type:uuid" json:"ptk_id,omitempty"`
	PesertaDidikID *string `gorm:"type:uuid" json:"peserta_didik_id,omitempty"`

	CreatedAt time.Time
	UpdatedAt time.Time

	// Relasi opsional
	RefreshTokens []RefreshToken `gorm:"foreignKey:PenggunaID"`
}
