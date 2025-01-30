package repository

import (
	"github.com/SthiraPs/DeliveryApp/db"
	"github.com/SthiraPs/DeliveryApp/models"
)

func CreateBranch(branch models.Branch) error {
	return db.DB.Create(&branch).Error
}

func GetBranches() ([]models.Branch, error) {
	var branches []models.Branch
	err := db.DB.Find(&branches).Error
	return branches, err
}
