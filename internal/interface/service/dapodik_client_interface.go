package service

import "school-backend/internal/dto"

type DapodikClientInterface interface {
	GetFullPesertaDidik(npsn string) (*struct {
		Rows []dto.SiswaDapodikFull `json:"rows"`
	}, error)
	GetFullPengguna(npsn string) (*struct {
		Rows []dto.PenggunaDapodikFull `json:"rows"`
	}, error)

	GetFullSekolah(baseUrl string, npsn string, token string) (*struct {
		Rows []dto.SekolahDapodikFull `json:"rows"`
	}, error)
}
