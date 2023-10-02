package models

import (
	"gorm.io/gorm"
)

type GoogleMapLocation struct {
	gorm.Model // This includes fields like ID, CreatedAt, UpdatedAt, and DeletedAt

	Name        string `gorm:"not null"`
	Description string
	Latitude    float64 `gorm:"not null"`
	Longitude   float64 `gorm:"not null"`
	UserID      uint
}
