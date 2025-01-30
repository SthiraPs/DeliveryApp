package models

import "time"

type ProductCategory struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StatusId    uint      `json:"statusId"`
	CreatedAt   time.Time `json:"createdAt" gorm:"not null" swaggerignore:"true"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"not null" swaggerignore:"true"`

	// Foreign keys
	Status Status `json:"status" gorm:"foreignKey:StatusId" swaggerignore:"true"`
}
