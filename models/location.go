package models

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model // This includes fields like ID, CreatedAt, UpdatedAt, and DeletedAt

	Name        string `gorm:"not null" binding:"required"`
	Description string
	Address     string `gorm:"not null" binding:"required"`
	Latitude    string
	Longitude   string
	UserID      uint
	Posts       []Post
}
