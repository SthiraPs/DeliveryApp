package repository

import (
	"github.com/SthiraPs/DeliveryApp/db"
	"github.com/SthiraPs/DeliveryApp/models"
)

func CreateDriverReview(driverReview models.DriverReview) error {
	return db.DB.Create(&driverReview).Error
}

func GetDriverReviews() ([]models.DriverReview, error) {
	var driverReviews []models.DriverReview
	err := db.DB.Find(&driverReviews).Error
	return driverReviews, err
}
