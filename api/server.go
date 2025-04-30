package api

import (
	"tumdum_backend/business"
	"tumdum_backend/config"
	"tumdum_backend/dao"
	"tumdum_backend/database"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	db     *gorm.DB
}

func NewServer() (*Server, error) {
	// Load configuration
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		return nil, err
	}

	// Initialize database
	db, err := database.InitDB(&cfg.Database)
	if err != nil {
		return nil, err
	}

	// Initialize router
	router := gin.Default()

	// Initialize DAO and services
	userDAO := dao.NewUserDAO(db)
	restaurantDAO := dao.NewRestaurantDAO(db)
	dishDAO := dao.NewDishDAO(db)
	orderDAO := dao.NewOrderDAO(db)

	userService := business.NewUserService(userDAO)
	restaurantService := business.NewRestaurantService(restaurantDAO)
	dishService := business.NewDishService(dishDAO)
	orderService := business.NewOrderService(orderDAO, dishDAO, restaurantDAO)

	// Initialize handlers
	userHandler := NewUserHandler(userService)
	restaurantHandler := NewRestaurantHandler(restaurantService)
	dishHandler := NewDishHandler(dishService)
	orderHandler := NewOrderHandler(orderService)

	// Setup routes
	api := router.Group("/api")
	{
		// User routes
		users := api.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUserByID)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
			users.GET("/:user_id/orders", orderHandler.GetUserOrders)
		}

		// Restaurant routes
		restaurants := api.Group("/restaurants")
		{
			restaurants.GET("", restaurantHandler.GetAllRestaurants)
			restaurants.GET("/:id", restaurantHandler.GetRestaurantByID)
			restaurants.GET("/:restaurant_id/dishes", dishHandler.GetDishesByRestaurantID)
		}

		// Dish routes
		dishes := api.Group("/dishes")
		{
			dishes.GET("/:id", dishHandler.GetDishByID)
		}

		// Order routes
		orders := api.Group("/orders")
		{
			orders.POST("", orderHandler.CreateOrder)
			orders.GET("/:id", orderHandler.GetOrderByID)
			orders.PUT("/:id/status", orderHandler.UpdateOrderStatus)
		}
	}

	// Setup Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return &Server{
		router: router,
		db:     db,
	}, nil
}

func (s *Server) Start(port string) error {
	return s.router.Run(":" + port)
}
