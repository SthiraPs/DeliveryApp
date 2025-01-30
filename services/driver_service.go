package services

import (
	"github.com/SthiraPs/DeliveryApp/models"
	"github.com/SthiraPs/DeliveryApp/repository"
	"github.com/SthiraPs/DeliveryApp/utils"
	"time"
)

func CreateDriverService(driver models.Driver) error {
	hashedPassword, err := utils.HashPassword(driver.Password)
	if err != nil {
		return err
	}

	driver.Password = hashedPassword
	driver.StatusId = 4 //Initially set to inactive
	driver.CreatedAt = time.Now()
	driver.UpdatedAt = time.Now()

	return repository.CreateDriverRepo(driver)
}

func GetAllDriversService() ([]models.Driver, error) {
	return repository.GetAllDriversRepo()
}

func GetDriverByIdService(id int) (models.Driver, error) {
	return repository.GetDriverByIdRepo(id)
}

func UpdateDriverByIdService(id int, updatedData models.Driver) (models.Driver, error) {
	return repository.UpdateDriverByIdRepo(id, updatedData)
}

func DeleteDriverByIdService(id int) (models.Driver, error) {
	return repository.DeleteDriverByIdRepo(id)
}
