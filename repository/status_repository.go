package repository

import (
	"github.com/SthiraPs/DeliveryApp/db"
	"github.com/SthiraPs/DeliveryApp/models"
)

func CreateStatus(status models.Status) error {
	return db.DB.Create(&status).Error
}

func GetStatuses() ([]models.Status, error) {
	var statuses []models.Status
	err := db.DB.Find(&statuses).Error
	return statuses, err
}
