package persistence

import (
	"school-backend/internal/entity"
	"school-backend/internal/interface/repository"

	"gorm.io/gorm"
)

type pembelajaranRepo struct {
	db *gorm.DB
}

func NewPembelajaranRepo(db *gorm.DB) repository.PembelajaranRepository {
	return &pembelajaranRepo{db: db}
}

func (r *pembelajaranRepo) FindAll() ([]*entity.Pembelajaran, error) {
	var pembelajaran []*entity.Pembelajaran
	err := r.db.Find(&pembelajaran).Error
	if err != nil {
		return nil, err
	}
	return pembelajaran, nil
}

func (r *pembelajaranRepo) Upsert(s *entity.Pembelajaran) error {
	return r.db.Save(s).Error
}
