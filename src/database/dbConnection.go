package database

import (
	"api/src/config"
	"api/src/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// OpenDbConnection returns the db connection opened
func OpenDbConnection() (db *gorm.DB, err error) {
	dsn := config.DbConnectionString
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	sqlDB, _ := db.DB()
	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		return nil, err
	}

	db.AutoMigrate(&models.User{})
	return db, err
}
