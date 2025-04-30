package dao

import (
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

func (dao *RestaurantDAO) GetAll() ([]models.Restaurant, error) {
	var restaurants []models.Restaurant
	err := dao.db.Find(&restaurants).Error
	return restaurants, err
}

func (dao *RestaurantDAO) Update(restaurant *models.Restaurant) error {
	return dao.db.Save(restaurant).Error
}

func (dao *RestaurantDAO) Delete(id uint) error {
	return dao.db.Delete(&models.Restaurant{}, id).Error
}
