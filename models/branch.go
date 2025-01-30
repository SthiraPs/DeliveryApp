package models

import "time"

type Branch struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	BranchName string    `json:"branchName"`
	PhoneNo    string    `json:"phoneNo"`
	Email      string    `json:"email"`
	StatusId   uint      `json:"statusId" gorm:"not null"`
	CreatedAt  time.Time `json:"createdAt" gorm:"not null" swaggerignore:"true"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"not null" swaggerignore:"true"`

	// Foreign keys
	Status Status `json:"status" gorm:"foreignKey:StatusId" swaggerignore:"true"`
}
