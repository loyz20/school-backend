package main

import (
	"log"
	"os"

	"school-backend/internal/entity"
	"school-backend/internal/infrastructure/persistence"
	"school-backend/internal/interface/controller"
	"school-backend/internal/usecase"
	"school-backend/pkg/jwt"
	"school-backend/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file tidak ditemukan, pakai environment langsung.")
	}
}

func main() {
	// Load env config (pastikan JWT_SECRET terisi)
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET belum diset di environment")
	}

	// Connect ke database
	dsn := os.Getenv("DATABASE_DSN") // contoh: "host=localhost user=postgres dbname=school sslmode=disable password=secret"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal konek DB: %v", err)
	}

	// Migrasi & seeder
	_ = db.AutoMigrate(&entity.Pengguna{}, &entity.RefreshToken{}) // dan entitas lainnya
	persistence.SeedDefaultUser(db)

	// Init repository
	userRepo := persistence.NewUserRepo(db)
	refreshRepo := persistence.NewRefreshTokenRepo(db)

	// Init JWT manager
	jwtManager := jwt.NewManager()

	// Init usecase
	authUsecase := usecase.NewAuthUsecase(refreshRepo, userRepo, jwtManager)

	// Init handler
	authHandler := controller.NewAuthHandler(authUsecase)

	// Setup Gin
	r := gin.Default()
	r.Use(middleware.Logger())

	// Auth endpoints
	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.RefreshToken)
	}

	authGroup := r.Group("/auth", middleware.AuthRequired())
	authGroup.POST("/logout-all", authHandler.LogoutAll)

	// Endpoint terlindungi
	api := r.Group("/api")
	api.Use(middleware.AuthRequired())
	{
		api.GET("/me", func(c *gin.Context) {
			userID := c.GetString("pengguna_id")
			role := c.GetString("peran_id_str")
			c.JSON(200, gin.H{"pengguna_id": userID, "peran": role})
		})
	}

	// Endpoint khusus admin
	admin := r.Group("/admin")
	admin.Use(middleware.AuthRequired(), middleware.RequireRoles("admin"))
	{
		admin.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Selamat datang admin!"})
		})
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server jalan di port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
