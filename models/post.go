package models

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Text     string  `json:"text" binding:"required"`
	UserID   uint    // default colum name will be user_id, you can specify it with `gorm:"column:desiredname"`
	Images   []Image // One-to-many relationship with images
	Comments []Comment
	// Other fields
}

type Image struct {
	gorm.Model
	Path   string `json:"url"` // Store the image path
	PostID uint   // Foreign key to associate the image with a post
}

var baseURL string
var staticImagePath string

// SetBaseURL sets the baseURL for the models package
func SetBaseURL(url string) {
	baseURL = url
}

func SetStaticImagePath(path string) {
	staticImagePath = path
}

func (i *Image) GetClickableURL() string {
	// Construct the full image URL by appending the path to the base URL
	return fmt.Sprintf("%s/%s/%s", baseURL, staticImagePath, filepath.Base(i.Path))
}

func (i *Image) MarshalJSON() ([]byte, error) {
	jsonMap := map[string]interface{}{
		"ID":        i.ID,
		"CreatedAt": i.CreatedAt,
		"UpdatedAt": i.UpdatedAt,
		"DeletedAt": i.DeletedAt,
		"url":       i.GetClickableURL(),
	}

	jsonString, err := json.Marshal(jsonMap)
	if err != nil {
		return nil, err
	}

	return jsonString, nil
}
