package repository

import (
	"github.com/SthiraPs/DeliveryApp/db"
	"github.com/SthiraPs/DeliveryApp/models"
)

func CreateProduct(product models.Product) error {
	return db.DB.Create(&product).Error
}

func GetProducts() ([]models.Product, error) {
	var products []models.Product
	err := db.DB.Find(&products).Error
	return products, err
}
