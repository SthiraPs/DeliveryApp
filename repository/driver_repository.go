package repository

import (
	"github.com/SthiraPs/DeliveryApp/db"
	"github.com/SthiraPs/DeliveryApp/models"
)

func CreateDriver(data models.Driver) error {
	return db.DB.Create(&data).Error
}

func GetAllDrivers() ([]models.Driver, error) {
	var drivers []models.Driver
	err := db.DB.Find(&drivers).Error
	return drivers, err
}

func GetDriverById(id int) (models.Driver, error) {
	var driver models.Driver
	err := db.DB.Where("id = ?", id).First(&driver).Error
	return driver, err
}

func UpdateDriverById(id int, updatedData models.Driver) (models.Driver, error) {
	var driver models.Driver
	err := db.DB.Where("id = ?", id).First(&driver).Error
	if err != nil {
		return models.Driver{}, err
	}

	err = db.DB.Model(&driver).Updates(updatedData).Error
	if err != nil {
		return models.Driver{}, err
	}
	return driver, nil
}

func DeleteDriverById(id int) (models.Driver, error) {
	var driver models.Driver
	err := db.DB.Where("id = ?", id).First(&driver).Error
	if err != nil {
		return models.Driver{}, err
	}

	driver.StatusId = 4

	err = db.DB.Model(&driver).Updates(driver).Error
	if err != nil {
		return models.Driver{}, err
	}
	return driver, nil
}
