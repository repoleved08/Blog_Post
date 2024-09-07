package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/repoleved08/blog/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func InitDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env %v",err.Error())
	}

	dsn := os.Getenv("DB_DSN")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to the database %v", err.Error())
	}
	
	fmt.Println("Database connected successfully!!")
	DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
}
