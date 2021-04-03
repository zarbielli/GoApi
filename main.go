package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnvVariables()
	fmt.Println("Running...")

	r := router.CreateRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
