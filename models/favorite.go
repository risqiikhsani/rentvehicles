package models

import (
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model // This includes fields like ID, CreatedAt, UpdatedAt, and DeletedAt

	UserID uint `validate:"required,numeric"`
	PostID uint `validate:"required,numeric"`
}
