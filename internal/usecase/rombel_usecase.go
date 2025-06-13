package usecase

import (
	"context"
	"school-backend/internal/entity"
	"school-backend/internal/interface/repository"
)

// RombelUsecase is the use case for managing Rombongan Belajar.
type RombelUsecase struct {
	repo repository.RombonganBelajarRepository
}

// NewRombelUsecase creates a new instance of RombelUsecase.
func NewRombelUsecase(rombelRepo repository.RombonganBelajarRepository) *RombelUsecase {
	return &RombelUsecase{
		repo: rombelRepo,
	}
}

func (uc *RombelUsecase) GetAll(ctx context.Context, page, limit int) ([]entity.RombonganBelajar, int, error) {
	offset := (page - 1) * limit
	return uc.repo.FindAll(ctx, limit, offset)
}
