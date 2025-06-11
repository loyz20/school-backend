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
	npsn := c.Query("npsn")
	token := c.Query("token")
	if npsn == "" && token == "" {
		response.Error(c, http.StatusBadRequest, "npsn dan token wajib diisi", nil)
		return
	}

	if _, err := h.usecase.ImportFromFullDapodikResponse(npsn); err != nil {
		response.Error(c, http.StatusInternalServerError, "Gagal import dari Dapodik", err.Error())
		return
	}

	response.Success(c, "Import berhasil", nil)
}
