package controller

import (
	"net/http"
	"school-backend/internal/usecase"
	"school-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type SekolahHandler struct {
	usecase *usecase.SekolahUsecase
}

func NewSekolahHandler(u *usecase.SekolahUsecase) *SekolahHandler {
	return &SekolahHandler{usecase: u}
}

func (h *SekolahHandler) GetAll(c *gin.Context) {
	result, err := h.usecase.GetAll(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Gagal mengambil data", nil)
		return
	}

	response.Success(c, "Data Sekolah", result)
}
