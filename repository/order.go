package repository

import (
	"github.com/SthiraPs/DeliveryApp/db"
	"github.com/SthiraPs/DeliveryApp/models"
)

func CreateOrder(order models.Order) error {
	return db.DB.Create(&order).Error
}

func GetOrders() ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Find(&orders).Error
	return orders, err
}
