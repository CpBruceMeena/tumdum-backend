package business

import (
	"errors"
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
	// Validate required fields
	if restaurant.Name == "" {
		return errors.New("restaurant name is required")
	}
	if restaurant.Email == "" {
		return errors.New("restaurant email is required")
	}
	if restaurant.Phone == "" {
		return errors.New("restaurant phone is required")
	}
	if restaurant.Address == "" {
		return errors.New("restaurant address is required")
	}
	if restaurant.City == "" {
		return errors.New("restaurant city is required")
	}
	if restaurant.State == "" {
		return errors.New("restaurant state is required")
	}
	if restaurant.Country == "" {
		return errors.New("restaurant country is required")
	}
	if restaurant.PostalCode == "" {
		return errors.New("restaurant postal code is required")
	}
	if restaurant.Cuisine == "" {
		return errors.New("restaurant cuisine is required")
	}

	// Validate rating
	if restaurant.Rating < 0 || restaurant.Rating > 5 {
		return errors.New("rating must be between 0 and 5")
	}

	return s.restaurantDAO.Create(restaurant)
}

func (s *RestaurantService) GetRestaurantByID(id uint) (*models.Restaurant, error) {
	return s.restaurantDAO.GetByID(id)
}

func (s *RestaurantService) GetAllRestaurants(cuisine, isActive, city string) ([]models.Restaurant, error) {
	return s.restaurantDAO.GetAll(cuisine, isActive, city)
}

func (s *RestaurantService) UpdateRestaurant(restaurant *models.Restaurant) error {
	// Validate required fields
	if restaurant.Name == "" {
		return errors.New("restaurant name is required")
	}
	if restaurant.Email == "" {
		return errors.New("restaurant email is required")
	}
	if restaurant.Phone == "" {
		return errors.New("restaurant phone is required")
	}
	if restaurant.Address == "" {
		return errors.New("restaurant address is required")
	}
	if restaurant.City == "" {
		return errors.New("restaurant city is required")
	}
	if restaurant.State == "" {
		return errors.New("restaurant state is required")
	}
	if restaurant.Country == "" {
		return errors.New("restaurant country is required")
	}
	if restaurant.PostalCode == "" {
		return errors.New("restaurant postal code is required")
	}
	if restaurant.Cuisine == "" {
		return errors.New("restaurant cuisine is required")
	}

	// Validate rating
	if restaurant.Rating < 0 || restaurant.Rating > 5 {
		return errors.New("rating must be between 0 and 5")
	}

	return s.restaurantDAO.Update(restaurant)
}

func (s *RestaurantService) DeleteRestaurant(id uint) error {
	return s.restaurantDAO.Delete(id)
}
