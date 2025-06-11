package usecase

import (
	"context"
	"school-backend/internal/dto"
	"school-backend/internal/entity"
	"school-backend/internal/interface/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) Update(ctx context.Context, u dto.UserUpdateRequest) error {
	user := entity.Pengguna{
		Nama:      u.Nama,
		Alamat:    u.Alamat,
		NoTelepon: u.NoTelepon,
		NoHP:      u.NoHP,
	}
	return uc.repo.Update(&user)
}

func (u *UserUsecase) GetAll(ctx context.Context, page, limit int) ([]entity.Pengguna, int, error) {
	offset := (page - 1) * limit
	return u.repo.FindAll(ctx, limit, offset)
}
