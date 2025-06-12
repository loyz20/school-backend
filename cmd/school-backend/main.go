package main

import (
	"log"

	"school-backend/internal/infrastructure/config"
	"school-backend/internal/infrastructure/persistence"
	"school-backend/internal/infrastructure/routes"
	"school-backend/pkg/utils"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := config.InitDB()

	// Migrasi & seede
	persistence.SeedDefaultUser(db)

	r := routes.SetupRouter(db)

	// Start server
	port := utils.GetEnv("PORT", "")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server jalan di port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
