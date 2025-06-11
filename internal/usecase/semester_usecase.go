package usecase

import (
	"context"
	"errors"
	"school-backend/internal/dto"
	"school-backend/internal/entity"
	"school-backend/internal/interface/repository"
	"school-backend/pkg/utils"
)

type SemesterUsecase struct {
	repo repository.SemesterRepository
}

func NewSemesterUsecase(repo repository.SemesterRepository) *SemesterUsecase {
	return &SemesterUsecase{repo: repo}
}

func (uc *SemesterUsecase) GetAll(ctx context.Context) ([]*entity.Semester, error) {
	return uc.repo.FindAll()
}

func (uc *SemesterUsecase) GetByID(ctx context.Context, id string) (*entity.Semester, error) {
	return uc.repo.FindByID(id)
}

func (uc *SemesterUsecase) Create(ctx context.Context, s dto.CreateSemesterRequest) (*entity.Semester, error) {
	// Contoh aturan bisnis: hanya boleh 1 semester aktif
	if s.Aktif {
		semesters, err := uc.repo.FindAll()
		if err != nil {
			return nil, err
		}
		for _, semester := range semesters {
			if semester.Aktif {
				return nil, errors.New("semester aktif sudah ada")
			}
		}
	}

	id, err := utils.GenerateSemesterID(s.TahunAjaran, s.SemesterKe)
	if err != nil {
		return nil, err
	}

	semester := &entity.Semester{
		SemesterID:  id,
		TahunAjaran: s.TahunAjaran,
		SemesterKe:  s.SemesterKe,
		Aktif:       s.Aktif,
	}

	err = uc.repo.Create(semester)
	if err != nil {
		return nil, err
	}
	return semester, nil
}

func (uc *SemesterUsecase) Update(ctx context.Context, s *entity.Semester) error {
	if s.Aktif {
		semesters, err := uc.repo.FindAll()
		if err != nil {
			return err
		}
		for _, semester := range semesters {
			if semester.Aktif && semester.SemesterID != s.SemesterID {
				return errors.New("semester aktif sudah ada")
			}
		}
	}
	return uc.repo.Update(s)
}

func (uc *SemesterUsecase) Delete(ctx context.Context, id string) error {
	return uc.repo.Delete(id)
}
