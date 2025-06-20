package usecase

import (
	"context"
	"school-backend/internal/entity"
	"school-backend/internal/interface/repository"
)

type SiswaUsecase struct {
	repo repository.SiswaRepository
}

func NewSiswaUsecase(repo repository.SiswaRepository) *SiswaUsecase {
	return &SiswaUsecase{repo: repo}
}

func (uc *SiswaUsecase) GetAll(ctx context.Context, page, limit int) ([]entity.Siswa, int, error) {
	offset := (page - 1) * limit
	return uc.repo.FindAll(ctx, limit, offset)
}

func (u *SiswaUsecase) GetByID(id string) (*entity.Siswa, error) {
	return u.repo.FindByID(id)
}

func (u *SiswaUsecase) Save(siswa *entity.Siswa) error {
	return u.repo.Upsert(siswa)
}

func (u *SiswaUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}
