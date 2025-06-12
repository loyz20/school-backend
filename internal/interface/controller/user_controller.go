package controller

import (
	"net/http"
	"school-backend/internal/entity"
	"school-backend/internal/usecase"
	"school-backend/pkg/response"
	"school-backend/pkg/validation"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

func (h *UserHandler) GetAll(c *gin.Context) {
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

	response.Paginated(c, "Data pengguna", result, pagination)
}

func (h *UserHandler) Update(ctx *gin.Context) {
	var input entity.Pengguna
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Validasi gagal", validation.FormatValidationError(err))
		return
	}

	User := h.usecase.Update(ctx, input)

	response.Created(ctx, "User berhasil dibuat", User)
}
