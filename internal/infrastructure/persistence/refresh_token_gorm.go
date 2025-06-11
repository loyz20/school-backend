package persistence

import (
	"school-backend/internal/entity"
	"school-backend/internal/interface/repository"
	"time"

	"gorm.io/gorm"
)

type refreshTokenRepo struct {
	db *gorm.DB
}

func NewRefreshTokenRepo(db *gorm.DB) repository.RefreshTokenRepository {
	return &refreshTokenRepo{db: db}
}

func (r *refreshTokenRepo) Save(token *entity.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *refreshTokenRepo) FindByToken(token string) (*entity.RefreshToken, error) {
	var rt entity.RefreshToken
	err := r.db.
		Where("token = ?", token).
		First(&rt).Error
	if err != nil {
		return nil, err
	}
	if rt.IsRevoked || rt.ExpiredAt.Before(time.Now()) {
		return nil, gorm.ErrRecordNotFound // atau custom error
	}
	return &rt, nil
}

func (r *refreshTokenRepo) RevokeByID(id string) error {
	return r.db.Model(&entity.RefreshToken{}).Where("id = ?", id).Update("is_revoked", true).Error
}

func (r *refreshTokenRepo) RevokeAllByUserID(userID string) error {
	return r.db.Model(&entity.RefreshToken{}).
		Where("pengguna_id = ? AND is_revoked = false", userID).
		Update("is_revoked", true).Error
}
