package api

import (
	"fmt"
	"tumdum_backend/business"
	"tumdum_backend/config"
	"tumdum_backend/middleware"

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
	imageHandler      *ImageHandler
}

// NewServer creates a new API server
func NewServer(
	restaurantService *business.RestaurantService,
	dishService *business.DishService,
	orderService *business.OrderService,
	userService *business.UserService,
	config *config.Config,
	imageHandler *ImageHandler,
) *Server {
	router := gin.Default()

	// Initialize handlers
	restaurantHandler := NewRestaurantHandler(restaurantService, imageHandler, dishService)
	dishHandler := NewDishHandler(dishService, imageHandler)
	orderHandler := NewOrderHandler(orderService)
	userHandler := NewUserHandler(userService)

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Create API group
	api := router.Group("/api")

	// Public routes
	userHandler.RegisterRoutes(api)
	imageHandler.RegisterRoutes(api)

	// Protected routes
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		restaurantHandler.RegisterRoutes(protected)
		dishHandler.RegisterRoutes(protected)
		orderHandler.RegisterRoutes(protected)
	}

	// Serve static files
	router.Static("/images", "./uploads")

	return &Server{
		router:            router,
		restaurantService: restaurantService,
		dishService:       dishService,
		orderService:      orderService,
		userService:       userService,
		config:            config,
		imageHandler:      imageHandler,
	}
}

// Start starts the server
func (s *Server) Start() error {
	return s.router.Run(fmt.Sprintf(":%d", s.config.Server.Port))
}
