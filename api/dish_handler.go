package api

import (
	"net/http"
	"strconv"
	"tumdum_backend/business"

	"github.com/gin-gonic/gin"
)

type DishHandler struct {
	dishService *business.DishService
}

func NewDishHandler(dishService *business.DishService) *DishHandler {
	return &DishHandler{dishService: dishService}
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
