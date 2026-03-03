package models

import "time"

type Customer struct {
	ID           int       `json:"id" db:"id"`
	Name         string    `json:"name" db:"name" validate:"required,max=50"`
	Phone        *string   `json:"phone" db:"phone"`
	Address      *string   `json:"address" db:"address"`
	City         *string   `json:"city" db:"city"`
	State        *string   `json:"state" db:"state"`
	Country      *string   `json:"country" db:"country"`
	ZipCode      *string   `json:"zip_code" db:"zip_code"`
	CreditRating *string   `json:"credit_rating" db:"credit_rating"`
	SalesRepID   *int      `json:"sales_rep_id" db:"sales_rep_id"`
	RegionID     *int      `json:"region_id" db:"region_id"`
	Comments     *string   `json:"comments" db:"comments"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	// Joined fields
	SalesRepName *string `json:"sales_rep_name,omitempty" db:"sales_rep_name"`
}

type CustomerFilter struct {
	Country    *string `json:"country"`
	SalesRepID *int    `json:"sales_rep_id"`
	SortBy     string  `json:"sort_by"` // "name", "phone", "sales_rep_id"
}

type CreateCustomerRequest struct {
	Name         string  `json:"name" validate:"required,max=50"`
	Phone        *string `json:"phone"`
	Address      *string `json:"address"`
	City         *string `json:"city"`
	State        *string `json:"state"`
	Country      *string `json:"country"`
	ZipCode      *string `json:"zip_code"`
	CreditRating *string `json:"credit_rating" validate:"omitempty,oneof=EXCELLENT GOOD POOR"`
	SalesRepID   *int    `json:"sales_rep_id"`
	RegionID     *int    `json:"region_id"`
	Comments     *string `json:"comments"`
}

type UpdateCustomerRequest struct {
	Name         *string `json:"name" validate:"omitempty,max=50"`
	Phone        *string `json:"phone"`
	Address      *string `json:"address"`
	City         *string `json:"city"`
	State        *string `json:"state"`
	Country      *string `json:"country"`
	ZipCode      *string `json:"zip_code"`
	CreditRating *string `json:"credit_rating" validate:"omitempty,oneof=EXCELLENT GOOD POOR"`
	SalesRepID   *int    `json:"sales_rep_id"`
	RegionID     *int    `json:"region_id"`
	Comments     *string `json:"comments"`
}
