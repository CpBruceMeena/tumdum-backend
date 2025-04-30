package api

import (
	"net/http"
	"strconv"
	"tumdum_backend/business"

	"github.com/gin-gonic/gin"
)

type RestaurantHandler struct {
	restaurantService *business.RestaurantService
}

func NewRestaurantHandler(restaurantService *business.RestaurantService) *RestaurantHandler {
	return &RestaurantHandler{restaurantService: restaurantService}
}

// @Summary Get all restaurants
// @Description Get a list of all restaurants
// @Tags restaurants
// @Produce json
// @Success 200 {array} models.Restaurant
// @Router /restaurants [get]
func (h *RestaurantHandler) GetAllRestaurants(c *gin.Context) {
	restaurants, err := h.restaurantService.GetAllRestaurants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, restaurants)
}

// @Summary Get restaurant by ID
// @Description Get restaurant details by ID
// @Tags restaurants
// @Produce json
// @Param id path int true "Restaurant ID"
// @Success 200 {object} models.Restaurant
// @Failure 404 {object} map[string]string
// @Router /restaurants/{id} [get]
func (h *RestaurantHandler) GetRestaurantByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	restaurant, err := h.restaurantService.GetRestaurantByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return
	}
	c.JSON(http.StatusOK, restaurant)
}
