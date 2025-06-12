package entity

type Semester struct {
	SemesterID  string `gorm:"primaryKey" json:"semester_id"`
	TahunAjaran string `json:"tahun_ajaran"` // Contoh: 2024/2025
	SemesterKe  int    `json:"semester_ke"`  // 1 = Ganjil, 2 = Genap
	Aktif       bool   `json:"aktif"`
}

func (Semester) TableName() string {
	return "semester"
}
