package dao

import (
	"tumdum_backend/models"

	"gorm.io/gorm"
)

type DishDAO struct {
	db *gorm.DB
}

func NewDishDAO(db *gorm.DB) *DishDAO {
	return &DishDAO{db: db}
}

func (dao *DishDAO) Create(dish *models.Dish) error {
	return dao.db.Create(dish).Error
}

func (dao *DishDAO) GetByID(id uint) (*models.Dish, error) {
	var dish models.Dish
	err := dao.db.Preload("Restaurant").First(&dish, id).Error
	if err != nil {
		return nil, err
	}
	return &dish, nil
}

func (dao *DishDAO) GetByRestaurantID(restaurantID uint) ([]models.Dish, error) {
	var dishes []models.Dish
	err := dao.db.Where("restaurant_id = ?", restaurantID).Find(&dishes).Error
	return dishes, err
}

func (dao *DishDAO) Update(dish *models.Dish) error {
	return dao.db.Save(dish).Error
}

func (dao *DishDAO) Delete(id uint) error {
	return dao.db.Delete(&models.Dish{}, id).Error
}
