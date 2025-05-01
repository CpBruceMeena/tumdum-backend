package models

import (
	"time"
)

type Dish struct {
	ID           uint       `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Price        float64    `json:"price"`
	RestaurantID uint       `json:"restaurant_id"`
	Restaurant   Restaurant `json:"restaurant" gorm:"foreignKey:RestaurantID"`
	Category     string     `json:"category"`
	IsAvailable  bool       `json:"is_available" gorm:"default:true"`
	ImageURL     string     `json:"image_url"`
}
