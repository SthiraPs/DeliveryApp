package services

import (
	"github.com/SthiraPs/DeliveryApp/models"
	"github.com/SthiraPs/DeliveryApp/repository"
)

func CreateUserService(user models.User) error {
	return repository.CreateUser(user)
}

func GetUsersService() ([]models.User, error) {
	return repository.GetUsers()
}
