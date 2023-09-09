package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Email    string `gorm:"uniqueIndex"`
	Name     string
	Phone    int32

	// Other fields
}
