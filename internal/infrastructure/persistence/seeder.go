package persistence

import (
	"log"
	"school-backend/internal/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedDefaultUser(db *gorm.DB) {
	const defaultUsername = "admin"

	var count int64
	db.Model(&entity.Pengguna{}).Where("username = ?", defaultUsername).Count(&count)
	if count > 0 {
		log.Println("[Seeder] User admin sudah ada, skip.")
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

	admin := entity.Pengguna{
		PenggunaID: "00000000-0000-0000-0000-000000000001", // UUID statis atau uuid.NewString()
		SekolahID:  "00000000-0000-0000-0000-000000000001", // UUID statis atau uuid.NewString()
		Nama:       "Administrator",
		Username:   defaultUsername,
		Password:   string(hashed),
		PeranIDStr: "admin",
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Printf("[Seeder] Gagal membuat user default: %v\n", err)
	} else {
		log.Println("[Seeder] User admin default berhasil dibuat: admin/admin123")
	}
}
