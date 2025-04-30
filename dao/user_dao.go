package dao

import (
	"tumdum_backend/models"

	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

func (dao *UserDAO) Create(user *models.User) error {
	return dao.db.Create(user).Error
}

func (dao *UserDAO) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := dao.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *UserDAO) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := dao.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *UserDAO) Update(user *models.User) error {
	return dao.db.Save(user).Error
}

func (dao *UserDAO) Delete(id uint) error {
	return dao.db.Delete(&models.User{}, id).Error
}
