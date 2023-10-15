package models

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Brand         string  `json:"brand" binding:"required"`
	BrandModel    string  `json:"brand_model" binding:"required"`
	VehicleType   string  `json:"vehicle_type" binding:"required"`
	Year          uint    `json:"year" binding:"required"`
	Transmission  string  `json:"transmission" binding:"required"`
	FuelType      string  `json:"fuel_type"`
	PricePerDay   uint    `json:"price_per_day" binding:"required"`
	PricePerWeek  uint    `json:"price_per_week" binding:"required"`
	PricePerMonth uint    `json:"price_per_month" binding:"required"`
	Discount      uint    `json:"discount"`
	Units         uint    `json:"units" binding:"required"`
	Available     bool    `json:"available" gorm:"default:true"`
	UserID        uint    // default colum name will be user_id, you can specify it with `gorm:"column:desiredname"`
	Images        []Image // One-to-many relationship with images
	LocationID    uint
	Reviews       []Review
	// Other fields
}

type Image struct {
	gorm.Model
	Path   string `json:"url"` // Store the image path
	PostID uint   // Foreign key to associate the image with a post
	RentID uint
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
