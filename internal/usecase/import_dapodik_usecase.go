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
	client            service.DapodikClientInterface
	semesterRepo      repository.SemesterRepository
	siswaRepo         repository.SiswaRepository
	userRepo          repository.UserRepository
	sekolahRepo       repository.SekolahRepository
	rombelRepo        repository.RombonganBelajarRepository
	anggotaRombelRepo repository.AnggotaRombelRepository
	pembelajaranRepo  repository.PembelajaranRepository
}

func NewImportDapodikUsecase(
	client service.DapodikClientInterface,
	semesterRepo repository.SemesterRepository,
	siswaRepo repository.SiswaRepository,
	userRepo repository.UserRepository,
	sekolahRepo repository.SekolahRepository,
	rombelRepo repository.RombonganBelajarRepository,
	anggotaRombelRepo repository.AnggotaRombelRepository,
	pembelajaranRepo repository.PembelajaranRepository,
) *ImportDapodikUsecase {
	return &ImportDapodikUsecase{
		client:            client,
		semesterRepo:      semesterRepo,
		siswaRepo:         siswaRepo,
		userRepo:          userRepo,
		sekolahRepo:       sekolahRepo,
		rombelRepo:        rombelRepo,
		anggotaRombelRepo: anggotaRombelRepo,
		pembelajaranRepo:  pembelajaranRepo,
	}
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

func (uc *ImportDapodikUsecase) ImportPengguna(npsn string) (*ImportResult, error) {
	resp, err := uc.client.GetFullPengguna(npsn)
	if err != nil {
		return nil, err
	}

	var result ImportResult

	for _, row := range resp.Rows {
		pengguna, err := row.ToEntity()
		if err != nil {
			result.JumlahGagal++
			continue
		}
		if err := uc.userRepo.Upsert(pengguna); err != nil {
			result.JumlahGagal++
			continue
		}
		result.JumlahBerhasil++
	}

	return &result, nil
}

func (uc *ImportDapodikUsecase) ImportSekolah(baseURL string, npsn string, token string) (*ImportResult, error) {
	resp, err := uc.client.GetFullSekolah(baseURL, npsn, token)
	if err != nil {
		return nil, err
	}

	var result ImportResult

	for _, row := range resp.Rows {
		sekolah, err := row.ToEntity(token)
		if err != nil {
			result.JumlahGagal++
			continue
		}
		if err := uc.sekolahRepo.Upsert(sekolah); err != nil {
			result.JumlahGagal++
			continue
		}
		result.JumlahBerhasil++
	}

	return &result, nil
}

func (uc *ImportDapodikUsecase) ImportRombel(npsn string) (*ImportResult, error) {
	resp, err := uc.client.GetFullRombel(npsn)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data dari dapodik: %w", err)
	}

	var result ImportResult

	for _, row := range resp.Rows {
		// Konversi rombel ke entitas domain
		rombel, err := row.ToEntity()
		if err != nil {
			result.JumlahGagal++
			continue
		}

		// Simpan rombel
		if err := uc.rombelRepo.Upsert(rombel); err != nil {
			result.JumlahGagal++
			continue
		}
		result.JumlahBerhasil++

		// Simpan anggota rombel (jika ada)
		for _, anggota := range row.AnggotaRombel {
			ent, err := anggota.ToEntity(row.RombonganBelajarID)
			if err != nil {
				result.JumlahGagal++
				continue
			}
			if err := uc.anggotaRombelRepo.Upsert(ent); err != nil {
				result.JumlahGagal++
				continue
			}
			result.JumlahBerhasil++
		}

		// Simpan pembelajaran (jika ada)
		for _, pb := range row.Pembelajaran {
			ent, err := pb.ToEntity(row.RombonganBelajarID)
			if err != nil {
				result.JumlahGagal++
				continue
			}
			if err := uc.pembelajaranRepo.Upsert(ent); err != nil {
				result.JumlahGagal++
				continue
			}
			result.JumlahBerhasil++
		}
	}

	return &result, nil
}
