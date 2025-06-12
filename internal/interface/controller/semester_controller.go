package controller

import (
	"net/http"
	"school-backend/internal/usecase"
	"school-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type SemesterHandler struct {
	usecase *usecase.SemesterUsecase
}

func NewSemesterHandler(u *usecase.SemesterUsecase) *SemesterHandler {
	return &SemesterHandler{usecase: u}
}

func (h *SemesterHandler) GetAll(c *gin.Context) {
	result, err := h.usecase.GetAll(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Gagal mengambil data", nil)
		return
	}

	response.Success(c, "Data Semester", result)
}
