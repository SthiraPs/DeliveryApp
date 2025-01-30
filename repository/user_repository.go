package repository

import (
	"github.com/SthiraPs/DeliveryApp/db"
	"github.com/SthiraPs/DeliveryApp/models"
)

func CreateUser(user models.User) error {
	return db.DB.Create(&user).Error
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	err := db.DB.Find(&users).Error
	return users, err
}
