package main

import (
	"log"
	"tumdum_backend/api"
	"tumdum_backend/business"
	"tumdum_backend/config"
	"tumdum_backend/dao"
	"tumdum_backend/database"
)

// @title Tumdum Backend API
// @version 1.0
// @description This is a sample server for Tumdum Backend.
// @host localhost:8080
// @BasePath /api
func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := database.InitDB(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize DAOs
	userDAO := dao.NewUserDAO(db)
	restaurantDAO := dao.NewRestaurantDAO(db)
	dishDAO := dao.NewDishDAO(db)
	orderDAO := dao.NewOrderDAO(db)

	// Initialize services
	userService := business.NewUserService(userDAO)
	restaurantService := business.NewRestaurantService(restaurantDAO)
	dishService := business.NewDishService(dishDAO)
	orderService := business.NewOrderService(orderDAO, dishDAO, restaurantDAO)

	// Initialize server
	server := api.NewServer(restaurantService, dishService, orderService, userService, cfg)

	// Start server
	log.Printf("Server starting on port %d...", cfg.Server.Port)
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
