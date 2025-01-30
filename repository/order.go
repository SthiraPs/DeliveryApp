package repository

import (
	"github.com/SthiraPs/DeliveryApp/db"
	"github.com/SthiraPs/DeliveryApp/models"
)

func CreateOrderRepo(order models.Order) error {
	tx := db.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Create the order (letting the DB assign an auto-increment ID)
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Assign the generated Order ID to each order item and reset their ID
	for i := range order.OrderItems {
		order.OrderItems[i].OrderId = order.ID
		order.OrderItems[i].ID = 0 // Reset ID to let DB generate it
	}

	// Create the order items (omit "ID" to prevent duplicate PK errors)
	if err := tx.Omit("ID").Create(&order.OrderItems).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}


func GetAllOrdersRepo() ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Preload("OrderItems").Find(&orders).Error
	return orders, err
}

func GetOrderByIdRepo(id int) (models.Order, error) {
	var order models.Order
	err := db.DB.Preload("OrderItems").Where("id = ?", id).First(&order).Error
	return order, err
}

func UpdateOrderByIdRepo(id int, updatedData models.Order) (models.Order, error) {
	var order models.Order
	err := db.DB.Where("id = ?", id).First(&order).Error
	if err != nil {
		return models.Order{}, err
	}

	// Update the Order fields
	err = db.DB.Model(&order).Updates(updatedData).Error
	if err != nil {
		return models.Order{}, err
	}

	for _, orderItem := range updatedData.OrderItems {
		err := db.DB.Where("id = ?", orderItem.ID).First(&orderItem).Error
		if err != nil {
			// If not found, create new OrderItem
			orderItem.OrderId = order.ID 
			err = db.DB.Create(&orderItem).Error
		} else {
			// If found, update OrderItem
			err = db.DB.Model(&orderItem).Updates(orderItem).Error
		}
		if err != nil {
			return models.Order{}, err
		}
	}

	return order, nil
}

func DeleteOrderByIdRepo(id int) (models.Order, error) {
	var order models.Order
	err := db.DB.Where("id = ?", id).First(&order).Error
	if err != nil {
		return models.Order{}, err
	}

	order.StatusId = 4

	err = db.DB.Model(&order).Updates(order).Error
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}
