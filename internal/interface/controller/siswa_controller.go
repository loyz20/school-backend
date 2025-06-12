package controller

import (
	"net/http"
	"school-backend/internal/usecase"
	"school-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type SiswaHandler struct {
	usecase *usecase.SiswaUsecase
}

func NewSiswaHandler(u *usecase.SiswaUsecase) *SiswaHandler {
	return &SiswaHandler{usecase: u}
}

func (h *SiswaHandler) GetAll(c *gin.Context) {
	page, limit := response.GetPaginationParams(c)

	result, total, err := h.usecase.GetAll(c.Request.Context(), page, limit)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Gagal mengambil data", nil)
		return
	}

	pagination := response.Pagination{
		Page:       page,
		Limit:      limit,
		TotalItems: total,
		TotalPages: response.CalculateTotalPages(total, limit),
	}

	response.Paginated(c, "Data siswa", result, pagination)
}
