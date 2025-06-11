package usecase

import (
	"errors"
	"school-backend/internal/entity"
	"school-backend/internal/interface/repository"
	"school-backend/pkg/auth"
	"school-backend/pkg/jwt"
	"time"
)

type AuthUsecase struct {
	userRepo   repository.UserRepository // Assuming you have a UserRepository for user operations
	repo       repository.RefreshTokenRepository
	jwtManager *jwt.Manager
}

func NewAuthUsecase(repo repository.RefreshTokenRepository, userRepo repository.UserRepository, jwtManager *jwt.Manager) *AuthUsecase {
	return &AuthUsecase{
		repo:       repo,
		userRepo:   userRepo,
		jwtManager: jwtManager,
	}

}

func (u *AuthUsecase) RefreshAccessToken(refreshTokenStr, ip, ua string) (string, string, error) {
	token, err := u.repo.FindByToken(refreshTokenStr)
	if err != nil || token.IsRevoked || token.ExpiredAt.Before(time.Now()) {
		return "", "", errors.New("refresh token tidak valid")
	}

	if token.IPAddress != ip || token.UserAgent != ua {
		_ = u.repo.RevokeByID(token.ID)
		return "", "", errors.New("kemungkinan token disalahgunakan")
	}

	// Revoke old token & create new token
	_ = u.repo.RevokeByID(token.ID)

	user, err := u.userRepo.FindByID(token.PenggunaID)
	if err != nil {
		return "", "", errors.New("akun tidak ditemukan")
	}

	accessToken, err := u.jwtManager.GenerateAccessToken(user.PenggunaID, user.PeranIDStr, ip, ua)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := u.jwtManager.GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}

	newToken := &entity.RefreshToken{
		PenggunaID: token.PenggunaID,
		Token:      newRefreshToken,
		IPAddress:  ip,
		UserAgent:  ua,
		ExpiredAt:  time.Now().AddDate(0, 0, 7),
	}

	if err := u.repo.Save(newToken); err != nil {
		return "", "", err
	}

	return accessToken, newRefreshToken, nil
}

func (u *AuthUsecase) Login(username, password, ip, ua string) (string, string, error) {
	// Ambil pengguna dari repo
	user, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return "", "", errors.New("akun tidak ditemukan")
	}

	// Verifikasi password (gunakan helper hash compare)
	if !auth.CheckPassword(password, user.Password) {
		return "", "", errors.New("password salah")
	}

	// Generate tokens
	accessToken, err := u.jwtManager.GenerateAccessToken(user.PenggunaID, user.PeranIDStr, ip, ua)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := u.jwtManager.GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}

	rt := &entity.RefreshToken{
		PenggunaID: user.PenggunaID,
		Token:      refreshToken,
		IPAddress:  ip,
		UserAgent:  ua,
		ExpiredAt:  time.Now().Add(7 * 24 * time.Hour),
	}

	if err := u.repo.Save(rt); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (u *AuthUsecase) LogoutAllDevices(userID string) error {
	return u.repo.RevokeAllByUserID(userID)
}
