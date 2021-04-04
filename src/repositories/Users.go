package repositories

import (
	"api/src/models"
	"fmt"

	"gorm.io/gorm"
)

// Users struct to represent users repository
type Users struct {
	db *gorm.DB
}

// UserRepository to create new user repository
func UserRepository(db *gorm.DB) *Users {
	return &Users{db}
}

// Create method isert a new user in db
func (repository Users) Create(user models.User) (uint32, error) {
	fmt.Println(repository.db.Create(user))

	return 0, nil
}
