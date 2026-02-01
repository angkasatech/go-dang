package main

import (
	"log"
	"net/http"
	"os"

	"go-dang/internal/category"
	"go-dang/internal/database"
	"go-dang/internal/router"
)

func main() {
	db := database.Connect()

	repo := category.NewRepository(db)
	service := category.NewService(repo)
	handler := category.NewHandler(service)

	r := router.Setup(handler)

	// Get port from environment variable, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	log.Printf("Server running on %s", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
