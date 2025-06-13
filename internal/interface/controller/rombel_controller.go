package controller

import (
	"net/http"
	"school-backend/internal/usecase"
	"school-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type RombelHandler struct {
	usecase *usecase.RombelUsecase
}

func NewRombelHandler(u *usecase.RombelUsecase) *RombelHandler {
	return &RombelHandler{usecase: u}
}

func (h *RombelHandler) GetAll(c *gin.Context) {
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

	response.Paginated(c, "Data rombongan belajar", result, pagination)
}
