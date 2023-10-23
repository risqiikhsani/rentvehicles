package models

import (
	"gorm.io/gorm"
)

type ForgotPassword struct {
	gorm.Model

	Token     string `gorm:"type:varchar(255);not null" validate:"required"`
	AccountID uint   `gorm:"not null" validate:"required,numeric"`
	Expired   bool   `gorm:"default:false" validate:"boolean"`
}
