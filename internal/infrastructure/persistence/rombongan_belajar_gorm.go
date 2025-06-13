package persistence

import (
	"context"
	"school-backend/internal/entity"
	"school-backend/internal/interface/repository"

	"gorm.io/gorm"
)

type rombelRepo struct {
	db *gorm.DB
}

func NewRombelRepo(db *gorm.DB) repository.RombonganBelajarRepository {
	return &rombelRepo{db}
}

func (r *rombelRepo) FindAll(ctx context.Context, limit, offset int) ([]entity.RombonganBelajar, int, error) {
	var users []entity.RombonganBelajar
	var count int64

	if err := r.db.WithContext(ctx).Model(&entity.RombonganBelajar{}).Count(&count).Error; err != nil {
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

func (r *rombelRepo) Upsert(s *entity.RombonganBelajar) error {
	return r.db.Save(s).Error
}
