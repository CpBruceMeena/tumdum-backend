package api

import (
	"net/http"
	"strconv"
	"tumdum_backend/business"
	"tumdum_backend/models"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService *business.OrderService
}

func NewOrderHandler(orderService *business.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

// @Summary Create a new order
// @Description Create a new order with the provided details
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.Order true "Order object"
// @Success 201 {object} models.Order
// @Failure 400 {object} map[string]string
// @Router /orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.orderService.CreateOrder(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// @Summary Get order by ID
// @Description Get order details by ID
// @Tags orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Failure 404 {object} map[string]string
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	order, err := h.orderService.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// @Summary Get user's order history
// @Description Get all orders for a specific user
// @Tags orders
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} models.Order
// @Failure 404 {object} map[string]string
// @Router /users/{user_id}/orders [get]
func (h *OrderHandler) GetUserOrders(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	orders, err := h.orderService.GetOrdersByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "orders not found"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// @Summary Update order status
// @Description Update the status of an order
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param status body string true "New status"
// @Success 200 {object} models.Order
// @Failure 400 {object} map[string]string
// @Router /orders/{id}/status [put]
func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var status struct {
		Status models.OrderStatus `json:"status"`
	}
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.orderService.UpdateOrderStatus(uint(id), status.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.orderService.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}
