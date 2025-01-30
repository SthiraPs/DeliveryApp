package repository

import (
	"errors"
	"github.com/SthiraPs/DeliveryApp/db"
	"github.com/SthiraPs/DeliveryApp/dto"
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

func GetAllUsersRepo() ([]dto.UserDTO, error) {
	var users []dto.UserDTO
	err := db.DB.
		Table("users").
		Select("users.id, users.first_name, users.last_name, users.email, users.phone_no, users.role_id, roles.role_name, users.status_id, statuses.status_name, users.created_at, users.updated_at").
		Joins("LEFT JOIN roles ON users.role_id = roles.id").
		Joins("LEFT JOIN statuses ON users.status_id = statuses.id").
		Find(&users).Error

	return users, err
}

func GetUserByIdRepo(id int) (dto.UserDTO, error) {
	var user dto.UserDTO
	err := db.DB.
		Table("users").
		Select("users.id, users.first_name, users.last_name, users.email, users.phone_no, users.role_id, roles.role_name, users.status_id, statuses.status_name, users.created_at, users.updated_at").
		Joins("LEFT JOIN roles ON users.role_id = roles.id").
		Joins("LEFT JOIN statuses ON users.status_id = statuses.id").
		First(&user).Where("users.id = ?", id).First(&user).Error
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
