package entity

import (
	"time"
)

type Siswa struct {
	RegistrasiID        string    `json:"registrasi_id"`
	NIPD                string    `gorm:"type:varchar(20)" json:"nipd"`
	PesertaDidikID      string    `gorm:"type:uuid" json:"peserta_didik_id"`
	Nama                string    `gorm:"type:varchar(100)" json:"nama"`
	NISN                string    `gorm:"type:varchar(20)" json:"nisn"`
	JenisKelamin        string    `gorm:"type:char(1)" json:"jenis_kelamin"`
	NIK                 string    `gorm:"type:varchar(20)" json:"nik"`
	TempatLahir         string    `gorm:"type:varchar(50)" json:"tempat_lahir"`
	TanggalLahir        time.Time `json:"tanggal_lahir"`
	AgamaIDStr          string    `gorm:"type:varchar(20)" json:"agama_id_str"`
	NomorTeleponRumah   string    `gorm:"type:varchar(20)" json:"nomor_telepon_rumah"`
	NomorTeleponSeluler string    `gorm:"type:varchar(20)" json:"nomor_telepon_seluler"`
	SemesterID          string    `gorm:"type:varchar(20)" json:"semester_id"`
	RombonganBelajarID  string    `gorm:"type:uuid" json:"rombongan_belajar_id"`
	NamaRombel          string    `gorm:"type:varchar(100)" json:"nama_rombel"`
	AnggotaRombelID     string    `gorm:"type:uuid" json:"anggota_rombel_id"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
