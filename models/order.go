package models

import "time"

type Order struct {
	ID          uint        `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	UserId      uint        `json:"userId"`
	TotalAmount float64     `json:"totalAmount" gorm:"type:decimal(10,2);"`
	StatusId    uint        `json:"statusId"`
	OrderItems  []OrderItem `json:"orderItems" gorm:"foreignKey:OrderId"`
	CreatedAt   time.Time   `json:"createdAt" gorm:"not null" swaggerignore:"true"`
	UpdatedAt   time.Time   `json:"updatedAt" gorm:"not null" swaggerignore:"true"`

	// Foreign keys
	User   User   `json:"user" gorm:"foreignKey:UserId" swaggerignore:"true"`
	Status Status `json:"status" gorm:"foreignKey:StatusId" swaggerignore:"true"`
}
