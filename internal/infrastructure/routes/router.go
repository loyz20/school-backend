package routes

import (
	"school-backend/internal/infrastructure/persistence"
	"school-backend/internal/interface/controller"
	"school-backend/internal/usecase"
	"school-backend/pkg/jwt"
	"school-backend/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.Logger())

	// Init repository
	userRepo := persistence.NewUserRepo(db)
	refreshRepo := persistence.NewRefreshTokenRepo(db)

	// Init JWT manager
	jwtManager := jwt.NewManager()

	// Init usecase
	authUsecase := usecase.NewAuthUsecase(refreshRepo, userRepo, jwtManager)

	// Init handler
	authHandler := controller.NewAuthHandler(authUsecase)

	api := r.Group("/api/v1")

	// Auth endpoints
	auth := api.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.RefreshToken)
		auth.POST("/logout", authHandler.LogoutAll, middleware.AuthRequired())
	}

	// Endpoint terlindungi
	api.Use(middleware.AuthRequired())
	{
		api.GET("/me", func(c *gin.Context) {
			userID := c.GetString("pengguna_id")
			role := c.GetString("peran_id_str")
			c.JSON(200, gin.H{"pengguna_id": userID, "peran": role})
		})
	}

	// Endpoint khusus admin
	admin := api.Group("/admin")
	admin.Use(middleware.AuthRequired(), middleware.RequireRoles("admin"))
	{
		admin.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Selamat datang admin!"})
		})
	}

	return r
}
