package models

import "time"

type Employee struct {
	ID            int        `json:"id" db:"id"`
	LastName      string     `json:"last_name" db:"last_name" validate:"required,max=25"`
	FirstName     *string    `json:"first_name" db:"first_name"`
	UserID        *string    `json:"userid" db:"userid"`
	StartDate     *time.Time `json:"start_date" db:"start_date"`
	Comments      *string    `json:"comments" db:"comments"`
	ManagerID     *int       `json:"manager_id" db:"manager_id"`
	Title         *string    `json:"title" db:"title"`
	DeptID        *int       `json:"dept_id" db:"dept_id"`
	Salary        *float64   `json:"salary" db:"salary"`
	CommissionPct *float64   `json:"commission_pct" db:"commission_pct"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
}

// SalesRep is a simplified employee view used for customer tree navigation.
type SalesRep struct {
	ID       int    `json:"id" db:"id"`
	FullName string `json:"full_name"`
	Title    string `json:"title" db:"title"`
}
