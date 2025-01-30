package models

import "time"

type DeliveryRequest struct {
	ID                 uint      `json:"id" gorm:"primaryKey"`
	UserId             uint       `json:"userId"`
	PickupLocation     string    `json:"pickupLocation"`
	DeliveryLocation   string    `json:"deliveryLocation"`
	PackageDescription string    `json:"packageDescription"`
	CostEst            float64   `json:"costEst" gorm:"type:decimal(10,2);"`
	DriverId            uint       `json:"riderId"`
	StatusId           uint       `json:"statusId"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null" swaggerignore:"true"`

	// Foreign keys
	Status Status `json:"status" gorm:"foreignKey:StatusId" swaggerignore:"true"`
	Driver Driver `json:"driver" gorm:"foreignKey:DriverId" swaggerignore:"true"`
}
