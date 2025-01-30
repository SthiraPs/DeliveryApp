package repository

import (
	"github.com/SthiraPs/DeliveryApp/db"
	"github.com/SthiraPs/DeliveryApp/models"
)

func CreateProductCategoryRepo(productCategory models.ProductCategory) error {
	return db.DB.Create(&productCategory).Error
}

func GetAllProductCategoriesRepo() ([]models.ProductCategory, error) {
	var productCategories []models.ProductCategory
	err := db.DB.Find(&productCategories).Error
	return productCategories, err
}

func GetProductCategoryByIdRepo(id int) (models.ProductCategory, error) {
	var productCategory models.ProductCategory
	err := db.DB.Where("id = ?", id).First(&productCategory).Error
	return productCategory, err
}

func UpdateProductCategoryByIdRepo(id int, updatedData models.ProductCategory) (models.ProductCategory, error) {
	var category models.ProductCategory
	err := db.DB.Where("id = ?", id).First(&category).Error
	if err != nil {
		return models.ProductCategory{}, err
	}

	err = db.DB.Model(&category).Updates(updatedData).Error
	if err != nil {
		return models.ProductCategory{}, err
	}
	return category, nil
}

func DeleteProductCategoryByIdRepo(id int) (models.ProductCategory, error) {
	var category models.ProductCategory
	err := db.DB.Where("id = ?", id).First(&category).Error
	if err != nil {
		return models.ProductCategory{}, err
	}

	category.StatusId = 4

	err = db.DB.Model(&category).Updates(category).Error
	if err != nil {
		return models.ProductCategory{}, err
	}
	return category, nil
}
