package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Path         string `json:"url"` // Store the image path
	PostID       *uint  // Foreign key to associate the image with a post
	RentDetailID *uint
	MainPostID   *uint
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

func (i Image) GetClickableURL() string {
	// Construct the full image URL by appending the path to the base URL
	return fmt.Sprintf("%s/%s/%s", baseURL, staticImagePath, filepath.Base(i.Path))
}

// https://pkg.go.dev/encoding/json
func (i Image) MarshalJSON() ([]byte, error) {
	jsonMap := map[string]interface{}{
		"ID":        i.ID,
		"CreatedAt": i.CreatedAt,
		"UpdatedAt": i.UpdatedAt,
		"DeletedAt": i.DeletedAt,
		"url":       i.GetClickableURL(),
	}

	return json.Marshal(jsonMap)

	// if err != nil {
	// 	return nil, err
	// }

	// return jsonString, nil
}

func (image *Image) BeforeDelete(tx *gorm.DB) (err error) {
	if image.Path != "" {
		if err := os.Remove(image.Path); err != nil {
			return err
		}
	}

	return nil
}
