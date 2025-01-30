package repository

import (
	"errors"
	"github.com/SthiraPs/DeliveryApp/db"
	"github.com/SthiraPs/DeliveryApp/models"
)

func SignInUserRepo(credentials models.User) (models.User, error) {
	var user models.User
	err := db.DB.Where("email = ?", credentials.Email).First(&user).Error
	if err != nil {
		return models.User{}, errors.New("invalid credentials. Please try again")
	}

	return user, nil // Successfully signed in
}

func CreateUserRepo(user models.User) error {
	return db.DB.Create(&user).Error
}

func GetAllUsersRepo() ([]models.User, error) {
	var users []models.User
	err := db.DB.Find(&users).Error
	return users, err
}

func GetUserByIdRepo(id int) (models.User, error) {
	var user models.User
	err := db.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func UpdateUserByIdRepo(id int, updatedData models.User) (models.User, error) {
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

func DeleteUserByIdRepo(id int) (models.User, error) {
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
