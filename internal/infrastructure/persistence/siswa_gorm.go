package persistence

import (
	"school-backend/internal/entity"
	"school-backend/internal/interface/repository"

	"gorm.io/gorm"
)

type siswaRepo struct {
	db *gorm.DB
}

func NewSiswaRepo(db *gorm.DB) repository.SiswaRepository {
	return &siswaRepo{db}
}

func (r *siswaRepo) FindAll() ([]*entity.Siswa, error) {
	var siswa []*entity.Siswa
	err := r.db.Find(&siswa).Error
	return siswa, err
}

func (r *siswaRepo) FindByID(id string) (*entity.Siswa, error) {
	var s entity.Siswa
	err := r.db.First(&s, "peserta_didik_id = ?", id).Error
	return &s, err
}

func (r *siswaRepo) Upsert(s *entity.Siswa) error {
	return r.db.Save(s).Error
}

func (r *siswaRepo) Delete(id string) error {
	return r.db.Delete(&entity.Siswa{}, "peserta_didik_id = ?", id).Error
}
