package business

import (
	"tumdum_backend/dao"
	"tumdum_backend/models"
)

type DishService struct {
	dishDAO *dao.DishDAO
}

func NewDishService(dishDAO *dao.DishDAO) *DishService {
	return &DishService{dishDAO: dishDAO}
}

func (s *DishService) CreateDish(dish *models.Dish) error {
	return s.dishDAO.Create(dish)
}

func (s *DishService) GetDishByID(id uint) (*models.Dish, error) {
	return s.dishDAO.GetByID(id)
}

func (s *DishService) GetDishesByRestaurantID(restaurantID uint) ([]models.Dish, error) {
	return s.dishDAO.GetByRestaurantID(restaurantID)
}

func (s *DishService) UpdateDish(dish *models.Dish) error {
	return s.dishDAO.Update(dish)
}

func (s *DishService) DeleteDish(id uint) error {
	return s.dishDAO.Delete(id)
}
