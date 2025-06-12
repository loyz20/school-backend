package usecase

import (
	"context"
	"school-backend/internal/entity"
	"school-backend/internal/interface/repository"
)

type SekolahUsecase struct {
	repo repository.SekolahRepository
}

func NewSekolahUsecase(repo repository.SekolahRepository) *SekolahUsecase {
	return &SekolahUsecase{repo: repo}
}

func (uc *SekolahUsecase) GetAll(ctx context.Context) ([]*entity.Sekolah, error) {
	return uc.repo.FindAll()
}
