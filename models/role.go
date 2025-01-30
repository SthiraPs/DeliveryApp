package models

import "time"

type Role struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	RoleName  string    `json:"roleName"`
	StatusId  uint      `json:"statusId"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null" swaggerignore:"true"`

	// Foreign keys
	Status Status `json:"status" gorm:"foreignKey:StatusId" swaggerignore:"true"`
}
