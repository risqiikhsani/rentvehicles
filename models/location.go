package models

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model // This includes fields like ID, CreatedAt, UpdatedAt, and DeletedAt

	Name        string `json:"name" form:"name"  validate:"required"`
	Description string `json:"description" form:"description"  validate:"required"`
	StreetName  string `json:"street_name" form:"street_name"  validate:"required"`
	Address     string `json:"address" form:"address"  validate:"required"`
	PostCode    string `json:"post_code" form:"post_code"  validate:"required"`
	City        string `json:"city" form:"city"  validate:"required"`
	Latitude    string `json:"latitude" form:"latitude" `
	Longitude   string `json:"longitude" form:"longitude"  `
	UserID      uint   `validate:"required,numeric"`
}
