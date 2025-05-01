package api

import (
	"fmt"
	"tumdum_backend/business"
	"tumdum_backend/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Server represents the API server
type Server struct {
	router            *gin.Engine
	db                *gorm.DB
	restaurantService *business.RestaurantService
	dishService       *business.DishService
	orderService      *business.OrderService
	userService       *business.UserService
	config            *config.Config
}

// NewServer creates a new API server
func NewServer(
	restaurantService *business.RestaurantService,
	dishService *business.DishService,
	orderService *business.OrderService,
	userService *business.UserService,
	config *config.Config,
) *Server {
	router := gin.Default()

	// Initialize handlers
	imageHandler := NewImageHandler("uploads")
	restaurantHandler := NewRestaurantHandler(restaurantService, imageHandler, dishService)
	dishHandler := NewDishHandler(dishService, imageHandler)
	orderHandler := NewOrderHandler(orderService)
	userHandler := NewUserHandler(userService)

	// Add CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Create API group
	api := router.Group("/api")

	// Register routes
	restaurantHandler.RegisterRoutes(api)
	dishHandler.RegisterRoutes(api)
	orderHandler.RegisterRoutes(api)
	userHandler.RegisterRoutes(api)
	imageHandler.RegisterRoutes(api)

	// Serve static files
	router.Static("/images", "./uploads")

	return &Server{
		router:            router,
		restaurantService: restaurantService,
		dishService:       dishService,
		orderService:      orderService,
		userService:       userService,
		config:            config,
	}
}

// Start starts the server
func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.config.Server.Port)
	return s.router.Run(addr)
}
