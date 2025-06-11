package repository

import "school-backend/internal/entity"

type UserRepository interface {
	FindByUsername(username string) (*entity.Pengguna, error)
	FindByID(id string) (*entity.Pengguna, error)
}
