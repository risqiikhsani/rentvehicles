package models

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model // This includes fields like ID, CreatedAt, UpdatedAt, and DeletedAt

	Name        string `gorm:"not null" validate:"required"`
	Description string
	Address     string `gorm:"not null" validate:"required"`
	Latitude    string
	Longitude   string
	UserID      uint `validate:"required,numeric"`
	Posts       []Post
}
