package repository

import (
	"context"
	"school-backend/internal/entity"
)

type SiswaRepository interface {
	FindAll(ctx context.Context, limit, offset int) ([]entity.Siswa, int, error)
	FindByID(id string) (*entity.Siswa, error)
	Upsert(siswa *entity.Siswa) error
	Delete(id string) error
}
