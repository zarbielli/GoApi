package database

import (
	"api/src/config"
	"api/src/models"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// OpenDbConnection returns the db connection opened
func OpenDbConnection() *gorm.DB {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := config.DbConnectionString
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	db.AutoMigrate(&models.User{})
	return db
}
