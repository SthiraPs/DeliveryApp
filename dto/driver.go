package dto

import "time"

type DriverDTO struct {
	ID            uint      `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	Email         string    `json:"email" gorm:"unique"`
	PhoneNo       string    `json:"phoneNo"`
	Password      string    `json:"password"`
	VehicleTypeId int       `json:"vehicleTypeId"`
	VehicleType   string    `json:"vehicleType"`
	VehicleRegNo  string    `json:"vehicleRegNo"`
	LicenceNo     string    `json:"licenceNo"`
	BranchId      int       `json:"branchId"`
	Branch        string    `json:"branch"`
	StatusId      int       `json:"statusId"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
