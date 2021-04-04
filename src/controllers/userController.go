package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateUser is responsable to create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error")
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal("Error")
	}

	db, err := database.OpenDbConnection()
	if err != nil {
		log.Fatal("error")
	}

	repository := repositories.UserRepository(db)
	repository.Create(user)
}

// GetAllUsers is responsable to recover all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting all users"))
}

// GetOneUser is responsible to recover a user by his ID
func GetOneUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a user"))
}

// UpdateUser is responsable to update one user using his ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating a user"))
}

// DeleteUser is responsable to delete one user,
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting a user"))
}
