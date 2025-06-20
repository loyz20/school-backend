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

	if _, err := h.usecase.ImportPesertaDidik(npsn); err != nil {
		response.Error(c, http.StatusInternalServerError, "Gagal import dari Dapodik", err.Error())
		return
	}

	response.Success(c, "Import berhasil", nil)
}

func (h *ImportDapodikHandler) ImportPG(c *gin.Context) {
	npsn := c.Query("npsn")
	token := c.Query("token")
	if npsn == "" && token == "" {
		response.Error(c, http.StatusBadRequest, "npsn dan token wajib diisi", nil)
		return
	}

	if _, err := h.usecase.ImportPengguna(npsn); err != nil {
		response.Error(c, http.StatusInternalServerError, "Gagal import dari Dapodik", err.Error())
		return
	}

	response.Success(c, "Import berhasil", nil)
}

func (h *ImportDapodikHandler) ImportSekolah(c *gin.Context) {
	npsn := c.Query("npsn")
	token := c.Query("token")
	baseUrl := c.Query("base_url")

	if npsn == "" && token == "" && baseUrl == "" {
		response.Error(c, http.StatusBadRequest, "npsn dan token dan url dapodik wajib diisi", nil)
		return
	}

	if _, err := h.usecase.ImportSekolah(baseUrl, npsn, token); err != nil {
		response.Error(c, http.StatusInternalServerError, "Gagal import dari Dapodik", err.Error())
		return
	}

	response.Success(c, "Import berhasil", nil)
}

func (h *ImportDapodikHandler) ImportSemester(c *gin.Context) {
	npsn := c.Query("npsn")
	token := c.Query("token")
	if npsn == "" && token == "" {
		response.Error(c, http.StatusBadRequest, "npsn dan token wajib diisi", nil)
		return
	}

	if err := h.usecase.ImportSemester(npsn); err != nil {
		response.Error(c, http.StatusInternalServerError, "Gagal import dari Dapodik", err.Error())
		return
	}

	response.Success(c, "Import berhasil", nil)
}

func (h *ImportDapodikHandler) ImportRombel(c *gin.Context) {
	npsn := c.Query("npsn")
	token := c.Query("token")
	if npsn == "" && token == "" {
		response.Error(c, http.StatusBadRequest, "npsn dan token wajib diisi", nil)
		return
	}

	if _, err := h.usecase.ImportRombel(npsn); err != nil {
		response.Error(c, http.StatusInternalServerError, "Gagal import dari Dapodik", err.Error())
		return
	}

	response.Success(c, "Import berhasil", nil)
}
