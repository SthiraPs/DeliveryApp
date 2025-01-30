package repository

import (
	"github.com/SthiraPs/DeliveryApp/db"
	"github.com/SthiraPs/DeliveryApp/models"
)

func CreateVehicleType(vehicleType models.VehicleType) error {
	return db.DB.Create(&vehicleType).Error
}

func GetCreateVehicleTypes() ([]models.VehicleType, error) {
	var vehicleTypes []models.VehicleType
	err := db.DB.Find(&vehicleTypes).Error
	return vehicleTypes, err
}
