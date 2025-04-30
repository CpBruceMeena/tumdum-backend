package models

import (
	"time"

	"gorm.io/gorm"
)

type Restaurant struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Address     string         `json:"address"`
	Phone       string         `json:"phone"`
	Rating      float32        `json:"rating"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
}
