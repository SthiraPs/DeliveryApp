package repository

import (
	"github.com/SthiraPs/DeliveryApp/db"
	"github.com/SthiraPs/DeliveryApp/models"
)

func CreateOrderItem(orderItem models.OrderItem) error {
	return db.DB.Create(&orderItem).Error
}

func GetOrderItems() ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	err := db.DB.Find(&orderItems).Error
	return orderItems, err
}
