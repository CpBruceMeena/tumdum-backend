package api

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"tumdum_backend/business"
	"tumdum_backend/models"

	"github.com/gin-gonic/gin"
)

type DishHandler struct {
	dishService  *business.DishService
	imageHandler *ImageHandler
}

func NewDishHandler(dishService *business.DishService, imageHandler *ImageHandler) *DishHandler {
	return &DishHandler{
		dishService:  dishService,
		imageHandler: imageHandler,
	}
}

// @Summary Create a new dish
// @Description Create a new dish with the provided details
// @Tags dishes
// @Accept multipart/form-data
// @Produce json
// @Param restaurant_id path int true "Restaurant ID"
// @Param name formData string true "Dish name"
// @Param description formData string true "Dish description"
// @Param price formData number true "Dish price"
// @Param category formData string true "Dish category"
// @Param image formData file false "Dish image"
// @Success 201 {object} models.Dish
// @Failure 400 {object} map[string]string
// @Router /restaurants/{restaurant_id}/dishes [post]
func (h *DishHandler) CreateDish(c *gin.Context) {
	restaurantID, err := strconv.ParseUint(c.Param("restaurant_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant id"})
		return
	}

	price, err := strconv.ParseFloat(c.PostForm("price"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid price"})
		return
	}

	dish := models.Dish{
		RestaurantID: uint(restaurantID),
		Name:         c.PostForm("name"),
		Description:  c.PostForm("description"),
		Price:        price,
		Category:     c.PostForm("category"),
	}

	// Handle image upload
	if imageFile, err := c.FormFile("image"); err == nil {
		imageURL, err := h.uploadImage(imageFile, uint(restaurantID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to upload image: %v", err)})
			return
		}
		dish.ImageURL = imageURL
	}

	if err := h.dishService.CreateDish(&dish); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dish)
}

// @Summary Update dish
// @Description Update dish details
// @Tags dishes
// @Accept multipart/form-data
// @Produce json
// @Param restaurant_id path int true "Restaurant ID"
// @Param dish_id path int true "Dish ID"
// @Param name formData string false "Dish name"
// @Param description formData string false "Dish description"
// @Param price formData number false "Dish price"
// @Param category formData string false "Dish category"
// @Param image formData file false "Dish image"
// @Success 200 {object} models.Dish
// @Failure 400 {object} map[string]string
// @Router /restaurants/{restaurant_id}/dishes/{dish_id} [put]
func (h *DishHandler) UpdateDish(c *gin.Context) {
	restaurantID, err := strconv.ParseUint(c.Param("restaurant_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant id"})
		return
	}

	dishID, err := strconv.ParseUint(c.Param("dish_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid dish id"})
		return
	}

	// Get existing dish
	dish, err := h.dishService.GetDishByID(uint(dishID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "dish not found"})
		return
	}

	// Verify that the dish belongs to the specified restaurant
	if dish.RestaurantID != uint(restaurantID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "dish does not belong to the specified restaurant"})
		return
	}

	// Update fields if provided
	if name := c.PostForm("name"); name != "" {
		dish.Name = name
	}
	if description := c.PostForm("description"); description != "" {
		dish.Description = description
	}
	if price := c.PostForm("price"); price != "" {
		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid price"})
			return
		}
		dish.Price = priceFloat
	}
	if category := c.PostForm("category"); category != "" {
		dish.Category = category
	}

	// Handle image upload
	if imageFile, err := c.FormFile("image"); err == nil {
		// Delete old image if exists
		if dish.ImageURL != "" {
			if err := h.imageHandler.DeleteImageByURL(dish.ImageURL); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to delete old image: %v", err)})
				return
			}
		}
		imageURL, err := h.uploadImage(imageFile, uint(dishID))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to upload image: %v", err)})
			return
		}
		dish.ImageURL = imageURL
	}

	if err := h.dishService.UpdateDish(dish); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dish)
}

// Helper function to upload image
func (h *DishHandler) uploadImage(file *multipart.FileHeader, dishID uint) (string, error) {
	// Create a temporary file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Create a new file in the uploads directory
	filename := fmt.Sprintf("dish_%d%s", dishID, filepath.Ext(file.Filename))
	uploadPath := filepath.Join("uploads", filename)

	// Create directory if it doesn't exist
	dir := filepath.Dir(uploadPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	dst, err := os.Create(uploadPath)
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

// @Summary Get dishes by restaurant ID
// @Description Get all dishes for a specific restaurant
// @Tags dishes
// @Produce json
// @Param restaurant_id path int true "Restaurant ID"
// @Success 200 {array} models.Dish
// @Failure 404 {object} map[string]string
// @Router /restaurants/{restaurant_id}/dishes [get]
func (h *DishHandler) GetDishesByRestaurantID(c *gin.Context) {
	restaurantID, err := strconv.ParseUint(c.Param("restaurant_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant id"})
		return
	}

	dishes, err := h.dishService.GetDishesByRestaurantID(uint(restaurantID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "dishes not found"})
		return
	}
	c.JSON(http.StatusOK, dishes)
}

// @Summary Get dish by ID
// @Description Get dish details by ID
// @Tags dishes
// @Produce json
// @Param id path int true "Dish ID"
// @Success 200 {object} models.Dish
// @Failure 404 {object} map[string]string
// @Router /dishes/{id} [get]
func (h *DishHandler) GetDishByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	dish, err := h.dishService.GetDishByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "dish not found"})
		return
	}
	c.JSON(http.StatusOK, dish)
}

// @Summary Get all dishes for a restaurant
// @Description Get all dishes for a specific restaurant
// @Tags dishes
// @Produce json
// @Param restaurant_id path int true "Restaurant ID"
// @Success 200 {array} models.Dish
// @Failure 400 {object} map[string]string
// @Router /restaurants/{restaurant_id}/dishes [get]
func (h *DishHandler) GetAllDishes(c *gin.Context) {
	restaurantID, err := strconv.ParseUint(c.Param("restaurant_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant id"})
		return
	}

	dishes, err := h.dishService.GetDishesByRestaurantID(uint(restaurantID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dishes)
}

// @Summary Delete dish
// @Description Delete a dish
// @Tags dishes
// @Produce json
// @Param restaurant_id path int true "Restaurant ID"
// @Param dish_id path int true "Dish ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /restaurants/{restaurant_id}/dishes/{dish_id} [delete]
func (h *DishHandler) DeleteDish(c *gin.Context) {
	restaurantID, err := strconv.ParseUint(c.Param("restaurant_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid restaurant id"})
		return
	}

	dishID, err := strconv.ParseUint(c.Param("dish_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid dish id"})
		return
	}

	// Get existing dish
	dish, err := h.dishService.GetDishByID(uint(dishID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "dish not found"})
		return
	}

	// Verify that the dish belongs to the specified restaurant
	if dish.RestaurantID != uint(restaurantID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "dish does not belong to the specified restaurant"})
		return
	}

	// Delete dish image if exists
	if dish.ImageURL != "" {
		if err := h.imageHandler.DeleteImageByURL(dish.ImageURL); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to delete dish image: %v", err)})
			return
		}
	}

	if err := h.dishService.DeleteDish(uint(dishID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dish deleted successfully"})
}

// RegisterRoutes registers the routes for the dish handler
func (h *DishHandler) RegisterRoutes(router *gin.RouterGroup) {
	dishes := router.Group("/restaurant-dishes")
	{
		dishes.POST("/:restaurant_id", h.CreateDish)
		dishes.GET("/:restaurant_id", h.GetDishesByRestaurantID)
		dishes.GET("/:restaurant_id/:dish_id", h.GetDishByID)
		dishes.PUT("/:restaurant_id/:dish_id", h.UpdateDish)
		dishes.DELETE("/:restaurant_id/:dish_id", h.DeleteDish)
	}
}
