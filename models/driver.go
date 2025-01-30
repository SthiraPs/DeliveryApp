package models

import "time"

type Driver struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	Email         string    `json:"email" gorm:"unique"`
	PhoneNo       string    `json:"phoneNo"`
	Password      string    `json:"password"`
	VehicleTypeId uint      `json:"vehicleTypeId"`
	VehicleRegNo  string    `json:"vehicleRegNo"`
	LicenceNo     string    `json:"licenceNo"`
	BranchId      uint      `json:"branchId"`
	StatusId      uint      `json:"statusId"`
	CreatedAt     time.Time `json:"createdAt" gorm:"not null" swaggerignore:"true"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"not null" swaggerignore:"true"`

	// Foreign keys
	VehicleType VehicleType `json:"vehicleType" gorm:"foreignKey:VehicleTypeId" swaggerignore:"true"`
	Branch      Branch      `json:"branch" gorm:"foreignKey:BranchId" swaggerignore:"true"`
	Status      Status      `json:"status" gorm:"foreignKey:StatusId" swaggerignore:"true"`
}
