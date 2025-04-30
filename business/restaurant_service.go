package business

import (
	"tumdum_backend/dao"
	"tumdum_backend/models"
)

type RestaurantService struct {
	restaurantDAO *dao.RestaurantDAO
}

func NewRestaurantService(restaurantDAO *dao.RestaurantDAO) *RestaurantService {
	return &RestaurantService{restaurantDAO: restaurantDAO}
}

func (s *RestaurantService) CreateRestaurant(restaurant *models.Restaurant) error {
	return s.restaurantDAO.Create(restaurant)
}

func (s *RestaurantService) GetRestaurantByID(id uint) (*models.Restaurant, error) {
	return s.restaurantDAO.GetByID(id)
}

func (s *RestaurantService) GetAllRestaurants() ([]models.Restaurant, error) {
	return s.restaurantDAO.GetAll()
}

func (s *RestaurantService) UpdateRestaurant(restaurant *models.Restaurant) error {
	return s.restaurantDAO.Update(restaurant)
}

func (s *RestaurantService) DeleteRestaurant(id uint) error {
	return s.restaurantDAO.Delete(id)
}
