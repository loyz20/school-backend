package entity

type Semester struct {
	SemesterID  string `json:"semester_id"`  // UUID
	TahunAjaran string `json:"tahun_ajaran"` // Contoh: 2024/2025
	SemesterKe  int    `json:"semester_ke"`  // 1 = Ganjil, 2 = Genap
	Aktif       bool   `json:"aktif"`

	RombonganBelajar []RombonganBelajar `gorm:"foreignKey:SemesterID"`
	Siswa            []Siswa            `gorm:"foreignKey:SemesterID"`
}
