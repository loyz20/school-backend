package repository

import (
	"context"
	"school-backend/internal/entity"
)

type RombonganBelajarRepository interface {
	FindAll(ctx context.Context, limit, offset int) ([]entity.RombonganBelajar, int, error)
	Upsert(s *entity.RombonganBelajar) error
}

type AnggotaRombelRepository interface {
	FindAll() ([]*entity.AnggotaRombel, error)
	Upsert(s *entity.AnggotaRombel) error
}

type PembelajaranRepository interface {
	FindAll() ([]*entity.Pembelajaran, error)
	Upsert(s *entity.Pembelajaran) error
}
