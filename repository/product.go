package repository

import (
	"github.com/SthiraPs/DeliveryApp/db"
	"github.com/SthiraPs/DeliveryApp/models"
)

func CreateProductRepo(product models.Product) error {
	return db.DB.Create(&product).Error
}

func GetAllProductsRepo() ([]models.Product, error) {
	var products []models.Product
	err := db.DB.Find(&products).Error
	return products, err
}

func GetProductByIdRepo(id int) (models.Product, error) {
	var product models.Product
	err := db.DB.Where("id = ?", id).First(&product).Error
	return product, err
}

func UpdateProductByIdRepo(id int, updatedData models.Product) (models.Product, error) {
	var product models.Product
	err := db.DB.Where("id = ?", id).First(&product).Error
	if err != nil {
		return models.Product{}, err
	}

	err = db.DB.Model(&product).Updates(updatedData).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func DeleteProductByIdRepo(id int) (models.Product, error) {
	var product models.Product
	err := db.DB.Where("id = ?", id).First(&product).Error
	if err != nil {
		return models.Product{}, err
	}

	product.StatusId = 4

	err = db.DB.Model(&product).Updates(product).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}
