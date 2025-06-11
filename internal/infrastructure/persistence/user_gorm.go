package persistence

import (
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
