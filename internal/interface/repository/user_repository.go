package repository

import (
	"context"
	"school-backend/internal/entity"
)

type UserRepository interface {
	FindByUsername(username string) (*entity.Pengguna, error)
	FindByID(id string) (*entity.Pengguna, error)
	FindAll(ctx context.Context, limit, offset int) ([]entity.Pengguna, int, error)
	Update(s *entity.Pengguna) error
	Upsert(s *entity.Pengguna) error
}
