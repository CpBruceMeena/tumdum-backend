package api

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"tumdum_backend/business"
	"tumdum_backend/models"

	"github.com/gin-gonic/gin"
)

type RestaurantHandler struct {
	restaurantService *business.RestaurantService
	imageHandler      *ImageHandler
	dishService       *business.DishService
}

func NewRestaurantHandler(restaurantService *business.RestaurantService, imageHandler *ImageHandler, dishService *business.DishService) *RestaurantHandler {
	return &RestaurantHandler{
		restaurantService: restaurantService,
		imageHandler:      imageHandler,
		dishService:       dishService,
	}
}

// @Summary Create restaurant
// @Description Create a new restaurant
// @Tags restaurants
// @Accept json
// @Produce json
// @Param restaurant body models.Restaurant true "Restaurant object"
// @Success 201 {object} models.Restaurant
// @Failure 400 {object} ErrorResponse
// @Router /restaurants [post]
func (h *RestaurantHandler) CreateRestaurant(c *gin.Context) {
	var restaurant models.Restaurant
	if err := c.ShouldBindJSON(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.restaurantService.CreateRestaurant(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, restaurant)
}

// @Summary Get all restaurants
// @Description Get a list of all restaurants with optional filtering
// @Tags restaurants
// @Accept json
// @Produce json
// @Param cuisine query string false "Filter by cuisine type"
// @Param rating query number false "Filter by minimum rating"
// @Param is_active query boolean false "Filter by active status"
// @Success 200 {array} models.Restaurant
// @Router /restaurants [get]
func (h *RestaurantHandler) GetAllRestaurants(c *gin.Context) {
	cuisine := c.Query("cuisine")
	isActive := c.Query("is_active")
	city := c.Query("city")

	restaurants, err := h.restaurantService.GetAllRestaurants(cuisine, isActive, city)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, restaurants)
}

// @Summary Get restaurant by ID
// @Description Get details of a specific restaurant
// @Tags restaurants
// @Accept json
// @Produce json
// @Param id path int true "Restaurant ID"
// @Success 200 {object} models.Restaurant
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /restaurants/{id} [get]
func (h *RestaurantHandler) GetRestaurantByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" || id == "undefined" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant ID"})
		return
	}

	restaurantID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant ID format"})
		return
	}

	restaurant, err := h.restaurantService.GetRestaurantByID(uint(restaurantID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return
	}

	c.JSON(http.StatusOK, restaurant)
}

// @Summary Update restaurant
// @Description Update an existing restaurant
// @Tags restaurants
// @Accept json
// @Produce json
// @Param id path int true "Restaurant ID"
// @Param restaurant body models.Restaurant true "Restaurant object"
// @Success 200 {object} models.Restaurant
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /restaurants/{id} [put]
func (h *RestaurantHandler) UpdateRestaurant(c *gin.Context) {
	id := c.Param("id")
	if id == "" || id == "undefined" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant ID"})
		return
	}

	restaurantID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant ID format"})
		return
	}

	var restaurant models.Restaurant
	if err := c.ShouldBindJSON(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	restaurant.ID = uint(restaurantID)
	if err := h.restaurantService.UpdateRestaurant(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, restaurant)
}

// Helper function to upload image
func (h *RestaurantHandler) uploadImage(file *multipart.FileHeader, imageType string) (string, error) {
	// Create a temporary file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Create a new file in the uploads directory
	filename := fmt.Sprintf("%s_%d%s", imageType, time.Now().UnixNano(), filepath.Ext(file.Filename))
	filepath := filepath.Join("uploads", filename)
	dst, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy the file
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return fmt.Sprintf("/images/%s", filename), nil
}

// Helper function to parse time
func parseTime(timeStr string) (time.Time, error) {
	return time.Parse("15:04:05", timeStr)
}

// @Summary Delete restaurant
// @Description Delete a restaurant
// @Tags restaurants
// @Accept json
// @Produce json
// @Param id path int true "Restaurant ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /restaurants/{id} [delete]
func (h *RestaurantHandler) DeleteRestaurant(c *gin.Context) {
	id := c.Param("id")
	if id == "" || id == "undefined" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant ID"})
		return
	}

	restaurantID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant ID format"})
		return
	}

	if err := h.restaurantService.DeleteRestaurant(uint(restaurantID)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "restaurant deleted successfully"})
}

// @Summary Get restaurant dishes
// @Description Get all dishes for a specific restaurant
// @Tags restaurants
// @Accept json
// @Produce json
// @Param id path int true "Restaurant ID"
// @Success 200 {array} models.Dish
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /restaurants/{id}/dishes [get]
func (h *RestaurantHandler) GetRestaurantDishes(c *gin.Context) {
	id := c.Param("id")
	if id == "" || id == "undefined" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant ID"})
		return
	}

	restaurantID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant ID format"})
		return
	}

	dishes, err := h.dishService.GetDishesByRestaurantID(uint(restaurantID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return
	}

	c.JSON(http.StatusOK, dishes)
}
