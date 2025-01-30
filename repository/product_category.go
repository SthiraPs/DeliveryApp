package repository

import (
	"github.com/SthiraPs/DeliveryApp/db"
	"github.com/SthiraPs/DeliveryApp/models"
)

func CreateProductCategory(productCategory models.ProductCategory) error {
	return db.DB.Create(&productCategory).Error
}

func GetProductCategories() ([]models.ProductCategory, error) {
	var productCategories []models.ProductCategory
	err := db.DB.Find(&productCategories).Error
	return productCategories, err
}
