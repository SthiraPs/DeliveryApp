package services

import (
	"errors"
	"github.com/SthiraPs/DeliveryApp/models"
	"github.com/SthiraPs/DeliveryApp/repository"
	"github.com/SthiraPs/DeliveryApp/utils"
	"time"
)

func SignInUserService(credentials models.User) (string, error) {
	user, err := repository.SignInUserRepo(credentials)
	if err != nil {
		return "", err
	}

	// Validate password
	if !utils.ValidatePassword(credentials.Password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.Email, user.FirstName, user.LastName, user.PhoneNo, user.RoleId, "user")
	if err != nil {
		return "", errors.New("error while generating token")
	}

	return token, nil
}

func CreateUserService(user models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return repository.CreateUserRepo(user)
}

func GetAllUsersService() ([]models.User, error) {
	return repository.GetAllUsersRepo()
}

func GetUserByIdService(id int) (models.User, error) {
	return repository.GetUserByIdRepo(id)
}

func UpdateUserByIdService(id int, updatedData models.User) (models.User, error) {
	return repository.UpdateUserByIdRepo(id, updatedData)
}

func DeleteUserByIdService(id int) (models.User, error) {
	return repository.DeleteUserByIdRepo(id)
}
