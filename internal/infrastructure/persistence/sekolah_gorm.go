package persistence

import (
	"school-backend/internal/entity"
	"school-backend/internal/interface/repository"

	"gorm.io/gorm"
)

type sekolahRepo struct {
	db *gorm.DB
}

func NewSekolahRepo(db *gorm.DB) repository.SekolahRepository {
	return &sekolahRepo{db: db}
}

func (r *sekolahRepo) FindAll() ([]*entity.Sekolah, error) {
	var sekolah []*entity.Sekolah
	err := r.db.Find(&sekolah).Error
	return sekolah, err
}

func (r *sekolahRepo) Upsert(s *entity.Sekolah) error {
	return r.db.Save(s).Error
}
