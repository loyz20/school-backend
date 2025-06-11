package persistence

import (
	"context"
	"school-backend/internal/entity"
	"school-backend/internal/interface/repository"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repository.UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) FindByUsername(username string) (*entity.Pengguna, error) {
	var user entity.Pengguna
	err := r.db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindByID(id string) (*entity.Pengguna, error) {
	var user entity.Pengguna
	err := r.db.First(&user, "pengguna_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) Update(s *entity.Pengguna) error {
	return r.db.Save(s).Error
}

func (r *userRepo) FindAll(ctx context.Context, limit, offset int) ([]entity.Pengguna, int, error) {
	var users []entity.Pengguna
	var count int64

	if err := r.db.WithContext(ctx).Model(&entity.Pengguna{}).Count(&count).Error; err != nil {
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
