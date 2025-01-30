package models

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price" gorm:"type:decimal(10,2);"`
	StockQty    int       `json:"stockQty"`
	CategoryId  uint      `json:"categoryId"`
	StatusId    uint      `json:"statusId"`
	CreatedAt   time.Time `json:"createdAt" gorm:"not null" swaggerignore:"true"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"not null" swaggerignore:"true"`

	// Foreign keys
	Category ProductCategory `json:"category" gorm:"foreignKey:CategoryId" swaggerignore:"true"`
	Status   Status          `json:"status" gorm:"foreignKey:StatusId" swaggerignore:"true"`
}
