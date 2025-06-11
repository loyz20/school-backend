package repository

import "school-backend/internal/entity"

type SemesterRepository interface {
	FindAll() ([]*entity.Semester, error)
	FindByID(id string) (*entity.Semester, error)
	Create(s *entity.Semester) error
	Update(s *entity.Semester) error
	Delete(id string) error
}
