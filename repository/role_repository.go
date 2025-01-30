package repository

import (
	"github.com/SthiraPs/DeliveryApp/db"
	"github.com/SthiraPs/DeliveryApp/models"
)

func CreateRole(role models.Role) error {
	return db.DB.Create(&role).Error
}

func GetRoles() ([]models.Role, error) {
	var roles []models.Role
	err := db.DB.Find(&roles).Error
	return roles, err
}
