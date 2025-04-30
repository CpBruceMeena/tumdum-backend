package business

import (
	"errors"
	"tumdum_backend/dao"
	"tumdum_backend/models"
)

type OrderService struct {
	orderDAO      *dao.OrderDAO
	dishDAO       *dao.DishDAO
	restaurantDAO *dao.RestaurantDAO
}

func NewOrderService(orderDAO *dao.OrderDAO, dishDAO *dao.DishDAO, restaurantDAO *dao.RestaurantDAO) *OrderService {
	return &OrderService{
		orderDAO:      orderDAO,
		dishDAO:       dishDAO,
		restaurantDAO: restaurantDAO,
	}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	// Validate restaurant exists
	restaurant, err := s.restaurantDAO.GetByID(order.RestaurantID)
	if err != nil || restaurant == nil {
		return errors.New("restaurant not found")
	}

	// Calculate total amount and validate dishes
	totalAmount := 0.0
	for i, item := range order.OrderItems {
		dish, err := s.dishDAO.GetByID(item.DishID)
		if err != nil || dish == nil {
			return errors.New("dish not found")
		}
		if dish.RestaurantID != order.RestaurantID {
			return errors.New("dish does not belong to the restaurant")
		}
		if !dish.IsAvailable {
			return errors.New("dish is not available")
		}
		order.OrderItems[i].Price = dish.Price
		totalAmount += dish.Price * float64(item.Quantity)
	}

	order.TotalAmount = totalAmount
	order.Status = models.OrderStatusPending

	return s.orderDAO.Create(order)
}

func (s *OrderService) GetOrderByID(id uint) (*models.Order, error) {
	return s.orderDAO.GetByID(id)
}

func (s *OrderService) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	return s.orderDAO.GetByUserID(userID)
}

func (s *OrderService) UpdateOrderStatus(orderID uint, status models.OrderStatus) error {
	order, err := s.orderDAO.GetByID(orderID)
	if err != nil {
		return err
	}

	// Validate status transition
	if !isValidStatusTransition(order.Status, status) {
		return errors.New("invalid status transition")
	}

	return s.orderDAO.UpdateStatus(orderID, status)
}

func isValidStatusTransition(current, new models.OrderStatus) bool {
	validTransitions := map[models.OrderStatus][]models.OrderStatus{
		models.OrderStatusPending:   {models.OrderStatusConfirmed, models.OrderStatusCancelled},
		models.OrderStatusConfirmed: {models.OrderStatusPreparing},
		models.OrderStatusPreparing: {models.OrderStatusReady},
		models.OrderStatusReady:     {models.OrderStatusDelivered},
	}

	for _, validStatus := range validTransitions[current] {
		if validStatus == new {
			return true
		}
	}
	return false
}
