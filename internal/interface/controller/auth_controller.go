package controller

import (
	"net/http"
	"school-backend/internal/dto"
	"school-backend/internal/usecase"
	"school-backend/pkg/response"
	"school-backend/pkg/validation"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	usecase *usecase.AuthUsecase
}

func NewAuthHandler(uc *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: uc}
}

// POST /auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validasi gagal", validation.FormatValidationError(err))
		return
	}

	ip := c.ClientIP()
	ua := c.GetHeader("User-Agent")

	accessToken, refreshToken, err := h.usecase.Login(req.Username, req.Password, ip, ua)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Login gagal", err.Error())
		return
	}

	response.Success(c, "Login berhasil", gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// POST /auth/refresh
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validasi gagal", validation.FormatValidationError(err))
		return
	}

	ip := c.ClientIP()
	ua := c.GetHeader("User-Agent")

	accessToken, refreshToken, err := h.usecase.RefreshAccessToken(req.RefreshToken, ip, ua)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Refresh token gagal", err.Error())
		return
	}

	response.Success(c, "Token berhasil diperbarui", gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (h *AuthHandler) LogoutAll(c *gin.Context) {
	penggunaID, ok := c.Get("pengguna_id")
	if !ok {
		response.Error(c, http.StatusUnauthorized, "pengguna_id tidak ditemukan", nil)
		return
	}

	err := h.usecase.LogoutAllDevices(penggunaID.(string))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Gagal logout semua sesi", err.Error())
		return
	}

	response.Success(c, "Berhasil logout dari semua perangkat", nil)
}
