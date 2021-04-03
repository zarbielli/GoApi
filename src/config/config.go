package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DbConnectionString string that set up db connection
	DbConnectionString string

	// Port Api running port
	Port int
)

// LoadEnvVariables fucntion responsable to load all env variables
func LoadEnvVariables() {
	var err error
	if err = godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	if Port, err = strconv.Atoi(os.Getenv("PORT")); err != nil {
		Port = 5000
	}

	DbConnectionString = fmt.Sprintf("host=db user=%s password=%s dbname=postgres port=9920 sslmode=disable TimeZone=America/SaoPaulo",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	)

}
