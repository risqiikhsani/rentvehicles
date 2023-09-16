package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Text   string  `json:"text" binding:"required"`
	UserID uint    // default colum name will be user_id, you can specify it with `gorm:"column:desiredname"`
	Images []Image // One-to-many relationship with images
	// Other fields
}

type Image struct {
	gorm.Model
	Path   string `json:"url"` // Store the image path
	PostID uint   // Foreign key to associate the image with a post
}
