package config

import (
	"fmt"
	"go-oauth2-gin/models"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDB() {
	var err error
	LoadEnv()

	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db

	// AutoMigrate your models
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	fmt.Println("📦 Database connected successfully")
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}