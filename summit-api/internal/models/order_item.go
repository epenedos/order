package models

import "time"

type OrderItem struct {
	OrdID           int       `json:"ord_id" db:"ord_id"`
	ItemID          int       `json:"item_id" db:"item_id"`
	ProductID       int       `json:"product_id" db:"product_id" validate:"required"`
	Price           *float64  `json:"price" db:"price"`
	Quantity        *int      `json:"quantity" db:"quantity" validate:"required,min=1"`
	QuantityShipped *int      `json:"quantity_shipped" db:"quantity_shipped"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	// Joined fields
	ProductName     *string `json:"product_name,omitempty"`
	ProductImageURL *string `json:"product_image_url,omitempty"`
}

type CreateOrderItemRequest struct {
	ProductID int  `json:"product_id" validate:"required"`
	Quantity  *int `json:"quantity" validate:"required,min=1"`
}

type UpdateOrderItemRequest struct {
	Quantity        *int `json:"quantity" validate:"omitempty,min=1"`
	QuantityShipped *int `json:"quantity_shipped"`
}
