package services

import (
	"github.com/SthiraPs/DeliveryApp/models"
	"github.com/SthiraPs/DeliveryApp/repository"
	"time"
)

func CreateProductCategoryService(productCategory models.ProductCategory) error {
	productCategory.CreatedAt = time.Now()
	productCategory.UpdatedAt = time.Now()

	return repository.CreateProductCategoryRepo(productCategory)
}

func GetAllProductCategoriesService() ([]models.ProductCategory, error) {
	return repository.GetAllProductCategoriesRepo()
}

func GetProductCategoryByIdService(id int) (models.ProductCategory, error) {
	return repository.GetProductCategoryByIdRepo(id)
}

func UpdateProductCategoryByIdService(id int, updatedData models.ProductCategory) (models.ProductCategory, error) {
	return repository.UpdateProductCategoryByIdRepo(id, updatedData)
}

func DeleteProductCategoryByIdService(id int) (models.ProductCategory, error) {
	return repository.DeleteProductCategoryByIdRepo(id)
}
