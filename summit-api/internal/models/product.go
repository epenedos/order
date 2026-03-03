package models

import "time"

type Product struct {
	ID                  int       `json:"id" db:"id"`
	Name                string    `json:"name" db:"name" validate:"required,max=50"`
	ShortDesc           *string   `json:"short_desc" db:"short_desc"`
	LongtextID          *int      `json:"longtext_id" db:"longtext_id"`
	ImageID             *int      `json:"image_id" db:"image_id"`
	SuggestedWhlslPrice *float64  `json:"suggested_whlsl_price" db:"suggested_whlsl_price"`
	WhlslUnits          *string   `json:"whlsl_units" db:"whlsl_units"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`
	// Joined/computed fields
	ImageURL    *string `json:"image_url,omitempty"`
	Description *string `json:"description,omitempty"`
}
