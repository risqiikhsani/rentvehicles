package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Text   string `json:"text" binding:"required"`
	UserID uint   // default colum name will be user_id, you can specify it with `gorm:"column:desiredname"`

	// Other fields
}
