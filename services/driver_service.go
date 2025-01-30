package services

import (
	"github.com/SthiraPs/DeliveryApp/models"
	"github.com/SthiraPs/DeliveryApp/repository"
)

func CreateDriverService(driver models.Driver) error {
	return repository.CreateDriver(driver)
}

func GetAllDriversService() ([]models.Driver, error) {
	return repository.GetAllDrivers()
}

func GetDriverByIdService(id int) (models.Driver, error) {
	return repository.GetDriverById(id)
}

func UpdateDriverByIdServer(id int, updatedData models.Driver) (models.Driver, error) {
	return repository.UpdateDriverById(id, updatedData)
}

func DeleteDriverByIdServer(id int) (models.Driver, error) {
	return repository.DeleteDriverById(id)
}
