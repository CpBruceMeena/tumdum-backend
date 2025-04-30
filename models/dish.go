package models

import (
	"time"

	"gorm.io/gorm"
)

type Dish struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Price        float64        `json:"price"`
	RestaurantID uint           `json:"restaurant_id"`
	Restaurant   Restaurant     `json:"restaurant" gorm:"foreignKey:RestaurantID"`
	IsAvailable  bool           `json:"is_available" gorm:"default:true"`
}
