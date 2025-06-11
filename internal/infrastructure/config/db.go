package config

import (
	"log"
	"school-backend/pkg/utils"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := utils.GetEnv("DATABASE_DSN", "") // contoh: "host=localhost user=postgres dbname=school sslmode=disable password=secret"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("❌ gagal koneksi database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ gagal dapatkan instance DB: %v", err)
	}

	// Optional: set connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ berhasil koneksi ke database")

	DB = db
	return db
}
