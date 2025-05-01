package models

import (
	"time"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "PENDING"
	OrderStatusConfirmed OrderStatus = "CONFIRMED"
	OrderStatusPreparing OrderStatus = "PREPARING"
	OrderStatusReady     OrderStatus = "READY"
	OrderStatusDelivered OrderStatus = "DELIVERED"
	OrderStatusCancelled OrderStatus = "CANCELLED"
)

type Order struct {
	ID           uint        `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	UserID       uint        `json:"user_id"`
	User         User        `json:"user" gorm:"foreignKey:UserID"`
	RestaurantID uint        `json:"restaurant_id"`
	Restaurant   Restaurant  `json:"restaurant" gorm:"foreignKey:RestaurantID"`
	Status       OrderStatus `json:"status" gorm:"type:varchar(20);default:'PENDING'"`
	TotalAmount  float64     `json:"total_amount"`
	OrderItems   []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	OrderID   uint      `json:"order_id"`
	Order     Order     `json:"-" gorm:"foreignKey:OrderID"`
	DishID    uint      `json:"dish_id"`
	Dish      Dish      `json:"dish" gorm:"foreignKey:DishID"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
}
