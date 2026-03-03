package models

import "time"

type Warehouse struct {
	ID        int       `json:"id" db:"id"`
	RegionID  int       `json:"region_id" db:"region_id"`
	Address   *string   `json:"address" db:"address"`
	City      *string   `json:"city" db:"city"`
	State     *string   `json:"state" db:"state"`
	Country   *string   `json:"country" db:"country"`
	ZipCode   *string   `json:"zip_code" db:"zip_code"`
	Phone     *string   `json:"phone" db:"phone"`
	ManagerID *int      `json:"manager_id" db:"manager_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
