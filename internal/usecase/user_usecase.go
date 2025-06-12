package usecase

import (
	"context"
	"school-backend/internal/entity"
	"school-backend/internal/interface/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) Update(ctx context.Context, u entity.Pengguna) error {
	return uc.repo.Update(&u)
}

func (uc *UserUsecase) GetAll(ctx context.Context, page, limit int) ([]entity.Pengguna, int, error) {
	offset := (page - 1) * limit
	return uc.repo.FindAll(ctx, limit, offset)
}
