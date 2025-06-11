package repository

import "school-backend/internal/entity"

type SiswaRepository interface {
	FindAll() ([]*entity.Siswa, error)
	FindByID(id string) (*entity.Siswa, error)
	Upsert(siswa *entity.Siswa) error
	Delete(id string) error
}
