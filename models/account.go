package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Username      string `gorm:"uniqueIndex" json:"username" binding:"required"`
	Email         string `gorm:"uniqueIndex" json:"email" binding:"required"`
	Password      string `gorm:"size:255;not null;" json:"password" binding:"required"`
	EmailVerified bool   `gorm:"default:false" json:"email_verified"`
	User          User   `json:"user"`
	// Other fields
}
