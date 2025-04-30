package business

import (
	"errors"
	"tumdum_backend/dao"
	"tumdum_backend/models"
)

type UserService struct {
	userDAO *dao.UserDAO
}

func NewUserService(userDAO *dao.UserDAO) *UserService {
	return &UserService{userDAO: userDAO}
}

func (s *UserService) CreateUser(user *models.User) error {
	// Check if user with same email exists
	existingUser, err := s.userDAO.GetByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("user with this email already exists")
	}

	return s.userDAO.Create(user)
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.userDAO.GetByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.userDAO.GetByEmail(email)
}

func (s *UserService) UpdateUser(user *models.User) error {
	existingUser, err := s.userDAO.GetByID(user.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}

	return s.userDAO.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.userDAO.Delete(id)
}
