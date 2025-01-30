package models

type VehicleType struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	TypeName string `json:"typeName"`
	StatusId uint   `json:"statusId" gorm:"not null"`

	// Foreign keys
	Status Status `json:"status" gorm:"foreignKey:StatusId" swaggerignore:"true"`
}
