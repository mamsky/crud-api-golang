package database

import (
	"crud/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

		var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek ke PostgreSQL:", err)
	}

	// Migrasi model
	err = DB.AutoMigrate(&models.Contact{})
	if err != nil {
		log.Fatal("Gagal migrasi model:", err)
	}

	log.Println("Sukses koneksi dan migrasi ke PostgreSQL")
}
