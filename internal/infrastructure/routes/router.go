package routes

import (
	"school-backend/internal/infrastructure/external"
	"school-backend/internal/infrastructure/persistence"
	"school-backend/internal/interface/controller"
	"school-backend/internal/usecase"
	"school-backend/pkg/jwt"
	"school-backend/pkg/middleware"
	"school-backend/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.Logger())

	// Init repository
	semesterRepo := persistence.NewSemesterRepo(db)
	userRepo := persistence.NewUserRepo(db)
	refreshRepo := persistence.NewRefreshTokenRepo(db)
	siswaRepo := persistence.NewSiswaRepo(db)
	sekolahRepo := persistence.NewSekolahRepo(db)
	// Init Rombongan Belajar repository
	rombelRepo := persistence.NewRombelRepo(db)
	// Init Anggota Rombel repository
	anggotaRombelRepo := persistence.NewAnggotaRombelRepo(db)
	// Init Pembelajaran repository
	pembelajaranRepo := persistence.NewPembelajaranRepo(db)

	// Init Dapodik client
	dapoUrl := utils.GetEnv("DAPO_URL", "localhost")
	dapoRepo := external.NewDapodikClient(dapoUrl)

	// Init JWT manager
	jwtManager := jwt.NewManager()

	// Init usecase
	authUsecase := usecase.NewAuthUsecase(refreshRepo, userRepo, jwtManager)
	dapoUseCase := usecase.NewImportDapodikUsecase(dapoRepo, semesterRepo, siswaRepo, userRepo, sekolahRepo, rombelRepo, anggotaRombelRepo, pembelajaranRepo)
	siswaUseCase := usecase.NewSiswaUsecase(siswaRepo)
	semesterUseCase := usecase.NewSemesterUsecase(semesterRepo)
	userUseCase := usecase.NewUserUsecase(userRepo)
	sekolahUseCase := usecase.NewSekolahUsecase(sekolahRepo)
	rombelUseCase := usecase.NewRombelUsecase(rombelRepo)

	// Init handler
	authHandler := controller.NewAuthHandler(authUsecase)
	dapoHandler := controller.NewAImportDapodikHandler(dapoUseCase)
	siswaHandler := controller.NewSiswaHandler(siswaUseCase)
	semesterHandler := controller.NewSemesterHandler(semesterUseCase)
	userHandler := controller.NewUserHandler(userUseCase)
	sekolahHandler := controller.NewSekolahHandler(sekolahUseCase)
	rombelHandler := controller.NewRombelHandler(rombelUseCase)

	api := r.Group("/api/v1")
	api.GET("/semester", semesterHandler.GetAll)
	api.GET("/sekolah", sekolahHandler.GetAll)

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
		api.GET("/siswa", siswaHandler.GetAll)
	}

	// Endpoint khusus admin
	admin := api.Group("/admin")
	admin.Use(middleware.AuthRequired(), middleware.RequireRoles("admin"))
	{
		admin.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Selamat datang admin!"})
		})
		admin.GET("/rombel", rombelHandler.GetAll)
		admin.GET("/sekolah", sekolahHandler.GetAll)
		admin.GET("/semester", semesterHandler.GetAll)
		admin.GET("/siswa", siswaHandler.GetAll)
		admin.GET("/user", userHandler.GetAll)
		admin.POST("/import/cek-koneksi", dapoHandler.ImportSekolah)
		admin.POST("/import/cek-semester", dapoHandler.ImportSemester)
		admin.POST("/import/siswa-dapodik", dapoHandler.ImportPD)
		admin.POST("/import/pengguna-dapodik", dapoHandler.ImportPG)
		admin.POST("/import/sekolah-dapodik", dapoHandler.ImportSekolah)
		admin.POST("/import/rombel-dapodik", dapoHandler.ImportRombel)
	}

	return r
}
