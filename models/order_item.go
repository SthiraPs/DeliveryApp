package models

type OrderItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	OrderId   uint    `json:"orderId"`
	ProductId uint    `json:"productId"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price" gorm:"type:decimal(10,2);"`

	// Foreign keys
	Order   Order   `json:"order" gorm:"foreignKey:OrderId;references:ID" swaggerignore:"true"`
	Product Product `json:"product" gorm:"foreignKey:ProductId;references:ID" swaggerignore:"true"`
}
