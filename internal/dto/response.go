package dto

type SemesterResponse struct {
	SemesterID  string `json:"id"`
	TahunAjaran string `json:"tahun_ajaran"`
	SemesterKe  int    `json:"semester_ke"`
	Aktif       bool   `json:"aktif"`
}

type UserResponse struct {
	Username       string  `json:"username"`
	Nama           string  `json:"nama"`
	PeranIDStr     string  `json:"peran_id_str"`
	Alamat         string  `json:"alamat"`
	NoTelepon      string  `json:"no_telepon"`
	NoHP           string  `json:"no_hp"`
	PtkID          *string `gorm:"type:uuid" json:"ptk_id,omitempty"`
	PesertaDidikID *string `gorm:"type:uuid" json:"peserta_didik_id,omitempty"`
}
