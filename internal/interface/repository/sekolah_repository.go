package repository

import (
	"school-backend/internal/entity"
)

type SekolahRepository interface {
	FindAll() ([]*entity.Sekolah, error)
	Upsert(s *entity.Sekolah) error
}
