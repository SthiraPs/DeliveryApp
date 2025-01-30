package repository

import (
	"github.com/SthiraPs/DeliveryApp/db"
	"github.com/SthiraPs/DeliveryApp/models"
)

func CreateDeliveryReqeust(deliveryRequest models.DeliveryRequest) error {
	return db.DB.Create(&deliveryRequest).Error
}

func GetDeliveryReqeusts() ([]models.DeliveryRequest, error) {
	var deliveryRequests []models.DeliveryRequest
	err := db.DB.Find(&deliveryRequests).Error
	return deliveryRequests, err
}
