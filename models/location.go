package models

import (
	"gorm.io/gorm"
)

type GoogleMapLocation struct {
	gorm.Model // This includes fields like ID, CreatedAt, UpdatedAt, and DeletedAt

	Name        string `gorm:"not null" binding:"required"`
	Description string
	Address     string  `binding:"required"`
	Latitude    float64 `gorm:"not null"`
	Longitude   float64 `gorm:"not null"`
	UserID      uint
}
