package dto

import "time"

type UserDTO struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	PhoneNo   string    `json:"phoneNo"`
	RoleId    int       `json:"roleId"`
	RoleName  string    `json:"roleName"`
	StatusId  int       `json:"statusID"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
