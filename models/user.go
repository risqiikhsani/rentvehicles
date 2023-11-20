package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PublicUsername string `gorm:"uniqueIndex" json:"publicUsername" validate:"required"`
	Name           string `json:"name" validate:"required"`
	About          string `json:"about"`
	Gender         string `json:"gender" validate:"omitempty,oneof=Male Female"`
	Role           string `gorm:"default:'Basic'" json:"role" validate:"omitempty,oneof=Admin Basic"`
	IsActive       string `gorm:"default:true" json:"is_active"`
	AccountID      uint   // default colum name will be account_id, you can specify it with `gorm:"column:desiredname"`

	// Other fields
}
