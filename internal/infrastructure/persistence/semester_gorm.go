package persistence

import (
	"school-backend/internal/entity"
	"school-backend/internal/interface/repository"

	"gorm.io/gorm"
)

type semesterRepo struct {
	db *gorm.DB
}

func NewSemesterRepo(db *gorm.DB) repository.SemesterRepository {
	return &semesterRepo{db: db}
}

func (r *semesterRepo) FindAll() ([]*entity.Semester, error) {
	var semesters []*entity.Semester
	err := r.db.Find(&semesters).Error
	return semesters, err
}

func (r *semesterRepo) FindByID(id string) (*entity.Semester, error) {
	var semester entity.Semester
	err := r.db.First(&semester, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &semester, nil
}

func (r *semesterRepo) Create(s *entity.Semester) error {
	return r.db.Create(s).Error
}

func (r *semesterRepo) Update(s *entity.Semester) error {
	return r.db.Save(s).Error
}

func (r *semesterRepo) Delete(id string) error {
	return r.db.Delete(&entity.Semester{}, "id = ?", id).Error
}
