package dao

import (
	"strconv"
	"tumdum_backend/models"

	"gorm.io/gorm"
)

type RestaurantDAO struct {
	db *gorm.DB
}

func NewRestaurantDAO(db *gorm.DB) *RestaurantDAO {
	return &RestaurantDAO{db: db}
}

func (dao *RestaurantDAO) Create(restaurant *models.Restaurant) error {
	return dao.db.Create(restaurant).Error
}

func (dao *RestaurantDAO) GetByID(id uint) (*models.Restaurant, error) {
	var restaurant models.Restaurant
	err := dao.db.First(&restaurant, id).Error
	if err != nil {
		return nil, err
	}
	return &restaurant, nil
}

func (dao *RestaurantDAO) GetAll(cuisine, isActive, city string) ([]models.Restaurant, error) {
	var restaurants []models.Restaurant
	query := dao.db.Model(&models.Restaurant{})

	// Apply filters if provided
	if cuisine != "" {
		query = query.Where("cuisine = ?", cuisine)
	}
	if isActive != "" {
		active, err := strconv.ParseBool(isActive)
		if err == nil {
			query = query.Where("is_active = ?", active)
		}
	}
	if city != "" {
		query = query.Where("city = ?", city)
	}

	err := query.Find(&restaurants).Error
	return restaurants, err
}

func (dao *RestaurantDAO) Update(restaurant *models.Restaurant) error {
	return dao.db.Save(restaurant).Error
}

func (dao *RestaurantDAO) Delete(id uint) error {
	return dao.db.Delete(&models.Restaurant{}, id).Error
}
