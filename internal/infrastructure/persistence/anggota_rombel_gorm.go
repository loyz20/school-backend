package persistence

import (
	"school-backend/internal/entity"
	"school-backend/internal/interface/repository"

	"gorm.io/gorm"
)

type anggotaRombelRepo struct {
	db *gorm.DB
}

func NewAnggotaRombelRepo(db *gorm.DB) repository.AnggotaRombelRepository {
	return &anggotaRombelRepo{db: db}
}

func (r *anggotaRombelRepo) FindAll() ([]*entity.AnggotaRombel, error) {
	var anggotaRombel []*entity.AnggotaRombel
	err := r.db.Find(&anggotaRombel).Error
	if err != nil {
		return nil, err
	}
	return anggotaRombel, nil
}

func (r *anggotaRombelRepo) Upsert(s *entity.AnggotaRombel) error {
	return r.db.Save(s).Error
}
