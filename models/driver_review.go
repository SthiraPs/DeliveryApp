package models

import "time"

type DriverReview struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Review    string    `json:"review"`
	Rating    float64   `json:"rating"`
	DriverId  uint       `json:"driverId"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null" swaggerignore:"true"`
}
