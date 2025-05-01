package models

import (
	"time"
)

type Restaurant struct {
	ID            uint      `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Email         string    `json:"email" gorm:"unique"`
	Phone         string    `json:"phone" gorm:"unique"`
	Address       string    `json:"address"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	Country       string    `json:"country"`
	PostalCode    string    `json:"postal_code"`
	Cuisine       string    `json:"cuisine"`
	Rating        float32   `json:"rating" gorm:"type:decimal(3,2);default:0.0"`
	IsActive      bool      `json:"is_active" gorm:"default:true"`
	LogoURL       string    `json:"logo_url"`
	CoverImageURL string    `json:"cover_image_url"`
	Dishes        []Dish    `json:"dishes,omitempty" gorm:"foreignKey:RestaurantID"`
	Orders        []Order   `json:"orders,omitempty" gorm:"foreignKey:RestaurantID"`
}
