package repository

import "school-backend/internal/entity"

type RefreshTokenRepository interface {
	Save(token *entity.RefreshToken) error
	FindByToken(token string) (*entity.RefreshToken, error)
	RevokeByID(id string) error
	RevokeAllByUserID(userID string) error
}
