package models

import "time"

type Image struct {
	ID          int       `json:"id" db:"id"`
	Format      *string   `json:"format" db:"format"`
	UseFilename bool      `json:"use_filename" db:"use_filename"`
	Filename    *string   `json:"filename" db:"filename"`
	ImageData   []byte    `json:"-" db:"image_data"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
