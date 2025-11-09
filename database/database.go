package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"news-feed-system-backend/app/models"
)

var DB *gorm.DB

func ConnectDB() {
    if err := godotenv.Load(".env"); err != nil {
		_ = godotenv.Load("../.env")
		_ = godotenv.Load("../../.env")
	}

    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Fatal("DATABASE_URL is not set. Make sure .env is loaded correctly.")
    }

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    DB = db

    if err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Follow{}); err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    log.Println("Database connected & migrated successfully")
}

