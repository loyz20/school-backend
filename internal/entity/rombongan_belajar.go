package entity

import "time"

type RombonganBelajar struct {
	RombonganBelajarID   string  `gorm:"column:rombongan_belajar_id;primaryKey;type:uuid"`
	Nama                 string  `gorm:"type:varchar(100)" json:"nama"`
	TingkatPendidikanID  string  `gorm:"type:varchar(10)"`
	TingkatPendidikanStr string  `gorm:"type:varchar(100)"`
	JenisRombel          string  `gorm:"type:varchar(5)" json:"jenis_rombel"`
	JenisRombelStr       string  `gorm:"type:varchar(100)"`
	KurikulumID          int     `gorm:"type:integer" json:"kurikulum_id"`
	KurikulumIDStr       string  `gorm:"type:varchar(100)"`
	IDRuang              string  `gorm:"type:uuid"`
	IDRuangStr           string  `gorm:"type:varchar(100)"`
	MovingClass          string  `gorm:"type:varchar(10)"`
	PTKID                string  `gorm:"type:uuid" json:"ptk_id"`
	PTKIDStr             string  `gorm:"column:ptk_id_str"`
	JurusanID            string  `gorm:"type:varchar(20)"`
	JurusanIDStr         string  `gorm:"type:varchar(100)"`
	IDKelasEkskul        *string `gorm:"type:uuid" json:"id_kelas_ekskul"`
	IDEkskul             *int    `gorm:"type:integer"`
	NMEkskul             *string `gorm:"type:varchar(100)"`
	SKEkskul             *string `gorm:"type:varchar(100)"`
	IDEkskulStr          *string `gorm:"type:varchar(100)"`

	SemesterID string   `gorm:"type:text" json:"semester_id"`
	Semester   Semester `gorm:"foreignKey:SemesterID;references:SemesterID"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoCreateTime" json:"updated_at"`

	AnggotaRombel []AnggotaRombel `gorm:"foreignKey:RombonganBelajarID"`
	Pembelajaran  []Pembelajaran  `gorm:"foreignKey:RombonganBelajarID"`
}

type AnggotaRombel struct {
	AnggotaRombelID     string `gorm:"column:anggota_rombel_id;primaryKey;type:uuid"`
	RombonganBelajarID  string `gorm:"type:uuid"`
	PesertaDidikID      string `gorm:"type:uuid"`
	JenisPendaftaranID  string `gorm:"type:varchar(5)"`
	JenisPendaftaranStr string `gorm:"type:varchar(100)"`
}

type Pembelajaran struct {
	PembelajaranID       string  `gorm:"column:pembelajaran_id;primaryKey;type:uuid"`
	RombonganBelajarID   string  `gorm:"type:uuid"`
	MataPelajaranID      int     `gorm:"type:integer"`
	MataPelajaranIDStr   string  `gorm:"column:matapelajaran_id_str"`
	NamaMataPelajaran    string  `gorm:"type:varchar(100)"`
	PTKTerdaftarID       string  `gorm:"type:uuid"`
	PTKID                string  `gorm:"type:uuid"`
	IndukPembelajaranID  *string `gorm:"type:uuid"`
	JamMengajarPerMinggu string  `gorm:"type:varchar(5)"`
	StatusDiKurikulum    string  `gorm:"type:varchar(5)"`
	StatusDiKurikulumStr string  `gorm:"column:status_kurikulum_str"`
}

func (RombonganBelajar) TableName() string {
	return "rombongan_belajar"
}

func (AnggotaRombel) TableName() string {
	return "anggota_rombel"
}

func (Pembelajaran) TableName() string {
	return "pembelajaran"
}
