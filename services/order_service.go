package services

import (
	"github.com/SthiraPs/DeliveryApp/models"
	"github.com/SthiraPs/DeliveryApp/repository"
	"time"
)

func CreateOrderService(order models.Order) error {
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	return repository.CreateOrderRepo(order)
}

func GetAllOrdersService() ([]models.Order, error) {
	return repository.GetAllOrdersRepo()
}

func GetOrderByIdService(id int) (models.Order, error) {
	return repository.GetOrderByIdRepo(id)
}

func UpdateOrderByIdService(id int, updatedData models.Order) (models.Order, error) {
	return repository.UpdateOrderByIdRepo(id, updatedData)
}

func DeleteOrderByIdService(id int) (models.Order, error) {
	return repository.DeleteOrderByIdRepo(id)
}
