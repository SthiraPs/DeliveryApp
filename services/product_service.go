package services

import (
	"github.com/SthiraPs/DeliveryApp/models"
	"github.com/SthiraPs/DeliveryApp/repository"
	"time"
)

func CreateProductService(product models.Product) error {
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	return repository.CreateProductRepo(product)
}

func GetAllProductsService() ([]models.Product, error) {
	return repository.GetAllProductsRepo()
}

func GetProductByIdService(id int) (models.Product, error) {
	return repository.GetProductByIdRepo(id)
}

func UpdateProductByIdService(id int, updatedData models.Product) (models.Product, error) {
	return repository.UpdateProductByIdRepo(id, updatedData)
}

func DeleteProductByIdService(id int) (models.Product, error) {
	return repository.DeleteProductByIdRepo(id)
}
