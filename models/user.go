package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"firstName" gorm:"not null"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email" gorm:"unique;not null"`
	PhoneNo   string    `json:"phoneNo" gorm:"unique;not null"`
	RoleId    uint      `json:"roleId" gorm:"not null"`
	Password  string    `json:"password" gorm:"not null"`
	StatusId  uint      `json:"statusId" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null" swaggerignore:"true"`

	// Foreign keys
	Role   Role   `json:"role" gorm:"foreignKey:RoleId" swaggerignore:"true"`
	Status Status `json:"status" gorm:"foreignKey:StatusId" swaggerignore:"true"`
}
