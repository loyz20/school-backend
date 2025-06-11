package dto

type SemesterResponse struct {
	SemesterID  string `json:"id"`
	TahunAjaran string `json:"tahun_ajaran"`
	SemesterKe  int    `json:"semester_ke"`
	Aktif       bool   `json:"aktif"`
}
