package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PublicUsername string `gorm:"uniqueIndex" json:"publicUsername" binding:"required"`
	Name           string `json:"name"`
	About          string `json:"about"`
	Gender         string `json:"gender"`
	Role           string `gorm:"default:'basic'" json:"role"`
	IsActive       string `gorm:"default:true" json:"is_active"`
	AccountID      uint   // default colum name will be account_id, you can specify it with `gorm:"column:desiredname"`

	// Other fields
}
