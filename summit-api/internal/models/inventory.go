package models

import "time"

type Inventory struct {
	ProductID              int        `json:"product_id" db:"product_id"`
	WarehouseID            int        `json:"warehouse_id" db:"warehouse_id"`
	AmountInStock          int        `json:"amount_in_stock" db:"amount_in_stock"`
	ReorderPoint           *int       `json:"reorder_point" db:"reorder_point"`
	MaxInStock             *int       `json:"max_in_stock" db:"max_in_stock"`
	OutOfStockExplanation  *string    `json:"out_of_stock_explanation" db:"out_of_stock_explanation"`
	RestockDate            *time.Time `json:"restock_date" db:"restock_date"`
	CreatedAt              time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at" db:"updated_at"`
	// Joined fields
	WarehouseCity *string `json:"warehouse_city,omitempty"`
	ProductName   *string `json:"product_name,omitempty"`
}
