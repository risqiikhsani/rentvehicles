package models

import (
	"gorm.io/gorm"
)

type Cat struct {
	gorm.Model
	Name   string `json:"name" form:"name" binding:"required" validate:"required,min=4,max=15"`
	Age    uint   `json:"age" form:"age" validate:"required,numeric,min=1"`
	Text   string `json:"text" form:"text" validate:"required"`
	UserID uint
	// Images []Image `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // One-to-many relationship with images
	// Other fields
}
