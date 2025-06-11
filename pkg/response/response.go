package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

func Created(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

func Error(ctx *gin.Context, statusCode int, message string, detail interface{}) {
	ctx.JSON(statusCode, gin.H{
		"status":  "error",
		"message": message,
		"errors":  detail,
	})
}

func BadRequest(ctx *gin.Context, detail interface{}) {
	Error(ctx, http.StatusBadRequest, "Permintaan tidak valid", detail)
}

func NotFound(ctx *gin.Context, detail interface{}) {
	Error(ctx, http.StatusNotFound, "Data tidak ditemukan", detail)
}

func InternalError(ctx *gin.Context, detail interface{}) {
	Error(ctx, http.StatusInternalServerError, "Terjadi kesalahan server", detail)
}
