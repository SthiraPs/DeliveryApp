package models

import "time"

type Status struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	StatusName string    `json:"statusName"`
	StatusType string    `json:"statusType"`
	CreatedAt  time.Time `json:"createdAt" gorm:"not null" swaggerignore:"true"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"not null" swaggerignore:"true"`
}
