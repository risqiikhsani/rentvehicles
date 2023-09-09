package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PublicUsername string `gorm:"uniqueIndex" json:"publicUsername" binding:"required"`
	Name           string `json:"name"`
	Phone          int32  `json:"phone"`
	About          string `json:"about"`
	Gender         string `json:"gender"`
	AccountID      uint

	// Other fields
}
