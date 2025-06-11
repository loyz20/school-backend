package usecase

import (
	"school-backend/internal/interface/repository"
	"school-backend/internal/interface/service"
)

type ImportResult struct {
	JumlahBerhasil int
	JumlahGagal    int
}

type ImportDapodikUsecase struct {
	client    service.DapodikClientInterface
	siswaRepo repository.SiswaRepository
}

func NewImportDapodikUsecase(client service.DapodikClientInterface, siswaRepo repository.SiswaRepository) *ImportDapodikUsecase {
	return &ImportDapodikUsecase{client: client, siswaRepo: siswaRepo}
}

func (uc *ImportDapodikUsecase) ImportFromFullDapodikResponse(sekolahID string) (*ImportResult, error) {
	resp, err := uc.client.GetFullPesertaDidik(sekolahID)
	if err != nil {
		return nil, err
	}

	var result ImportResult

	for _, row := range resp.Rows {
		siswa, err := row.ToEntity()
		if err != nil {
			result.JumlahGagal++
			continue
		}
		if err := uc.siswaRepo.Upsert(siswa); err != nil {
			result.JumlahGagal++
			continue
		}
		result.JumlahBerhasil++
	}

	return &result, nil
}
