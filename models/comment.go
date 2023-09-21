package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Text   string `json:"text" binding:"required"`
	UserID uint   `binding:"required"` // default colum name will be user_id, you can specify it with `gorm:"column:desiredname"`
	PostID uint   `binding:"required"`
	// Other fields
}
