package usecase

import (
	"errors"
	"fmt"
	"school-backend/internal/entity"
	"school-backend/internal/interface/repository"
	"school-backend/internal/interface/service"
	"school-backend/pkg/utils"

	"gorm.io/gorm"
)

type ImportResult struct {
	JumlahBerhasil int
	JumlahGagal    int
}

type ImportDapodikUsecase struct {
	client       service.DapodikClientInterface
	semesterRepo repository.SemesterRepository
	siswaRepo    repository.SiswaRepository
}

func NewImportDapodikUsecase(client service.DapodikClientInterface, semesterRepo repository.SemesterRepository, siswaRepo repository.SiswaRepository) *ImportDapodikUsecase {
	return &ImportDapodikUsecase{client: client, semesterRepo: semesterRepo, siswaRepo: siswaRepo}
}

func (uc *ImportDapodikUsecase) ImportPesertaDidik(npsn string) (*ImportResult, error) {
	resp, err := uc.client.GetFullPesertaDidik(npsn)
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

func (uc *ImportDapodikUsecase) ImportSemester(npsn string) error {
	resp, err := uc.client.GetFullPesertaDidik(npsn)
	if err != nil {
		return err
	}

	// Ambil hanya 1 entri pertama untuk dapatkan semester_id
	if len(resp.Rows) == 0 || resp.Rows[0].SemesterID == "" {
		return errors.New("tidak ada data peserta didik atau semester_id kosong")
	}

	row := resp.Rows[0]

	_, err = uc.semesterRepo.FindByID(row.SemesterID)
	if err == nil {
		// Semester sudah ada, tidak perlu ditambahkan
		return nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err // error DB lain
	}

	// Semester belum ada, tambahkan baru
	tahun, semester, err := utils.ParseSemesterID(row.SemesterID)
	if err != nil {
		return fmt.Errorf("gagal parse semester_id: %w", err)
	}

	smt := &entity.Semester{
		SemesterID:  row.SemesterID,
		TahunAjaran: tahun,
		SemesterKe:  semester,
		Aktif:       true,
	}
	return uc.semesterRepo.Create(smt)
}
