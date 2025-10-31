package main

import (
	"log"
	"net/http"

	"github.com/multidownloader/backend/internal/api"
)

func main() {
	r := api.SetupRoutes()

	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
