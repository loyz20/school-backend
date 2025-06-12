package entity

import "time"

type Pengguna struct {
	PenggunaID     string  `gorm:"primaryKey;type:uuid" json:"pengguna_id"`
	SekolahID      string  `gorm:"type:uuid" json:"sekolah_id"`
	Username       string  `json:"username"`
	Nama           string  `json:"nama"`
	PeranIDStr     string  `json:"peran_id_str"`
	Password       string  `json:"password"`
	Alamat         *string `json:"alamat"`
	NoTelepon      *string `json:"no_telepon"`
	NoHP           *string `json:"no_hp"`
	PTKID          *string `json:"ptk_id"`
	PesertaDidikID *string `json:"peserta_didik_id"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoCreateTime" json:"updated_at"`
}

func (Pengguna) TableName() string {
	return "pengguna"
}
