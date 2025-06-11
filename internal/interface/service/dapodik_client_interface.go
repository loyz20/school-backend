package service

import "school-backend/internal/dto"

type DapodikClientInterface interface {
	GetFullPesertaDidik(sekolahID string) (*struct {
		Rows []dto.SiswaDapodikFull `json:"rows"`
	}, error)
}
