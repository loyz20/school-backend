package persistence

import (
	"context"
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

func (r *siswaRepo) FindAll(ctx context.Context, limit, offset int) ([]entity.Siswa, int, error) {
	var users []entity.Siswa
	var count int64

	if err := r.db.WithContext(ctx).Model(&entity.Siswa{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Order("created_at desc").
		Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, int(count), nil
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
