package business

import (
	"errors"
	"strconv"
	"strings"
	"tumdum_backend/auth"
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
	err := s.userDAO.Create(user)
	if err != nil {
		if strings.Contains(err.Error(), "users_email_unique") {
			return errors.New("user with this email already exists")
		}
		return err
	}
	return nil
}

func (s *UserService) GetUser(id string) (*models.User, error) {
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}
	return s.userDAO.GetByID(uint(userID))
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.userDAO.GetByEmail(email)
}

func (s *UserService) UpdateUser(id string, user *models.User) error {
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return errors.New("invalid user ID")
	}
	existingUser, err := s.userDAO.GetByID(uint(userID))
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}

	// Don't allow updating email
	user.Email = existingUser.Email
	// Don't allow updating password through this method
	user.Password = existingUser.Password
	user.ID = uint(userID)

	return s.userDAO.Update(user)
}

func (s *UserService) DeleteUser(id string) error {
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return errors.New("invalid user ID")
	}
	return s.userDAO.Delete(uint(userID))
}

func (s *UserService) UpdatePassword(id string, currentPassword, newPassword string) error {
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return errors.New("invalid user ID")
	}
	user, err := s.userDAO.GetByID(uint(userID))
	if err != nil {
		return err
	}

	// Verify current password
	if !auth.CheckPasswordHash(currentPassword, user.Password) {
		return errors.New("current password is incorrect")
	}

	// Hash new password
	hashedPassword, err := auth.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// Update password
	user.Password = hashedPassword
	return s.userDAO.Update(user)
}
