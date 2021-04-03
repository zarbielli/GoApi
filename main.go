package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running...")

	r := router.CreateRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
}
