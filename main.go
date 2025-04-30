package main

import (
	"log"
	"tumdum_backend/api"
)

// @title Tumdum Backend API
// @version 1.0
// @description This is a sample server for Tumdum Backend.
// @host localhost:8080
// @BasePath /api
func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	log.Println("Server starting on port 8080...")
	if err := server.Start("8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 