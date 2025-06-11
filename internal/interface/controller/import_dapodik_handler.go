package controller

import (
	"net/http"
	"school-backend/internal/usecase"
	"school-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type ImportDapodikHandler struct {
	usecase *usecase.ImportDapodikUsecase
}

func NewAImportDapodikHandler(uc *usecase.ImportDapodikUsecase) *ImportDapodikHandler {
	return &ImportDapodikHandler{usecase: uc}
}

func (h *ImportDapodikHandler) ImportPD(c *gin.Context) {
	sekolahID := c.Query("sekolah_id")
	if sekolahID == "" {
		response.Error(c, http.StatusBadRequest, "sekolah_id wajib diisi", nil)
		return
	}

	if _, err := h.usecase.ImportFromFullDapodikResponse(sekolahID); err != nil {
		response.Error(c, http.StatusInternalServerError, "Gagal import dari Dapodik", err.Error())
		return
	}

	response.Success(c, "Import berhasil", nil)
}
