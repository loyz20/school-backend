package main

import (
	"log"
	"os"

	"school-backend/internal/entity"
	"school-backend/internal/infrastructure/config"
	"school-backend/internal/infrastructure/persistence"
	"school-backend/internal/infrastructure/routes"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := config.InitDB()

	// Migrasi & seeder
	_ = db.AutoMigrate(&entity.Pengguna{}, &entity.RefreshToken{}) // dan entitas lainnya
	persistence.SeedDefaultUser(db)

	r := routes.SetupRouter(db)

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
