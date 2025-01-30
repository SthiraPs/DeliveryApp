package services

import (
	"github.com/SthiraPs/DeliveryApp/models"
	"github.com/SthiraPs/DeliveryApp/repository"
)

func CreateUserService(user models.User) error {
	return repository.CreateUser(user)
}

func GetAllUsersService() ([]models.User, error) {
	return repository.GetAllUsers()
}

func GetUserByIdService(id int) (models.User, error) {
	return repository.GetUserById(id)
}

func UpdateUserByIdServer(id int, updatedData models.User) (models.User, error) {
	return repository.UpdateUserById(id, updatedData)
}

func DeleteUserByIdServer(id int) (models.User, error) {
	return repository.DeleteUserById(id)
}
