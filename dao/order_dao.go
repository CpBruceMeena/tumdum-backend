package dao

import (
	"tumdum_backend/models"

	"gorm.io/gorm"
)

type OrderDAO struct {
	db *gorm.DB
}

func NewOrderDAO(db *gorm.DB) *OrderDAO {
	return &OrderDAO{db: db}
}

func (dao *OrderDAO) Create(order *models.Order) error {
	return dao.db.Create(order).Error
}

func (dao *OrderDAO) GetByID(id uint) (*models.Order, error) {
	var order models.Order
	err := dao.db.Preload("User").Preload("Restaurant").Preload("OrderItems.Dish").
		First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (dao *OrderDAO) GetByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := dao.db.Preload("Restaurant").Preload("OrderItems.Dish").
		Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (dao *OrderDAO) UpdateStatus(orderID uint, status models.OrderStatus) error {
	return dao.db.Model(&models.Order{}).Where("id = ?", orderID).
		Update("status", status).Error
}

func (dao *OrderDAO) Update(order *models.Order) error {
	return dao.db.Save(order).Error
}
