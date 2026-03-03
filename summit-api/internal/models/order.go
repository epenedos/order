package models

import "time"

type Order struct {
	ID           int        `json:"id" db:"id"`
	CustomerID   int        `json:"customer_id" db:"customer_id" validate:"required"`
	DateOrdered  *time.Time `json:"date_ordered" db:"date_ordered"`
	DateShipped  *time.Time `json:"date_shipped" db:"date_shipped"`
	SalesRepID   *int       `json:"sales_rep_id" db:"sales_rep_id"`
	Total        *float64   `json:"total" db:"total"`
	PaymentType  *string    `json:"payment_type" db:"payment_type" validate:"omitempty,oneof=CASH CREDIT"`
	OrderFilled  bool       `json:"order_filled" db:"order_filled"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	// Joined fields
	CustomerName *string     `json:"customer_name,omitempty" db:"customer_name"`
	SalesRepName *string     `json:"sales_rep_name,omitempty" db:"sales_rep_name"`
	Items        []OrderItem `json:"items,omitempty"`
}

type CreateOrderRequest struct {
	CustomerID  int        `json:"customer_id" validate:"required"`
	DateOrdered *time.Time `json:"date_ordered"`
	SalesRepID  *int       `json:"sales_rep_id"`
	PaymentType *string    `json:"payment_type" validate:"omitempty,oneof=CASH CREDIT"`
}

type UpdateOrderRequest struct {
	DateOrdered *time.Time `json:"date_ordered"`
	DateShipped *time.Time `json:"date_shipped"`
	SalesRepID  *int       `json:"sales_rep_id"`
	PaymentType *string    `json:"payment_type" validate:"omitempty,oneof=CASH CREDIT"`
	OrderFilled *bool      `json:"order_filled"`
}
