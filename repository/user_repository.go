package repository

import (
	"github.com/SthiraPs/DeliveryApp/db"
	"github.com/SthiraPs/DeliveryApp/models"
)

func CreateUser(user models.User) error {
	return db.DB.Create(&user).Error
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := db.DB.Find(&users).Error
	return users, err
}

func GetUserById(id int) (models.User, error) {
	var user models.User
	err := db.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func UpdateUserById(id int, updatedData models.User) (models.User, error) {
	var user models.User
	err := db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	err = db.DB.Model(&user).Updates(updatedData).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func DeleteUserById(id int) (models.User, error) {
	var user models.User
	err := db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	user.StatusId = 4

	err = db.DB.Model(&user).Updates(user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
