package main

import (
	"log"
	"net/http"

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

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
