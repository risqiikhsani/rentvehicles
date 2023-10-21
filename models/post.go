package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Brand         string   `json:"brand" form:"brand" binding:"required"`
	BrandModel    string   `json:"brand_model" form:"brand_model" binding:"required"`
	VehicleType   string   `json:"vehicle_type" form:"vehicle_type" binding:"required"`
	Year          uint     `json:"year" form:"year" binding:"required"`
	Transmission  string   `json:"transmission" form:"transmission" binding:"required"`
	FuelType      string   `json:"fuel_type" form:"fuel_type" binding:"required"`
	PricePerDay   uint     `json:"price_per_day" form:"price_per_day" binding:"required"`
	PricePerWeek  uint     `json:"price_per_week" form:"price_per_week" binding:"required"`
	PricePerMonth uint     `json:"price_per_month" form:"price_per_month" binding:"required"`
	Discount      uint     `json:"discount" form:"discount"`
	Units         uint     `json:"units" form:"units" binding:"required"`
	Available     bool     `json:"available" form:"available" gorm:"default:true"`
	UserID        uint     // default colum name will be user_id, you can specify it with `gorm:"column:desiredname"`
	Images        []Image  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // One-to-many relationship with images
	LocationID    uint     `json:"location_id" form:"location_id" binding:"required"`
	Reviews       []Review `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Other fields
}

type Image struct {
	gorm.Model
	Path   string `json:"url"` // Store the image path
	PostID *uint  // Foreign key to associate the image with a post
	RentID *uint
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

func (post *Post) AfterDelete(tx *gorm.DB) (err error) {
	// First, fetch all associated images
	var images []Image
	tx.Model(post).Association("Images").Find(&images)

	// Delete each associated image
	for _, image := range images {
		tx.Unscoped().Delete(&image)
	}

	return
}

func (image *Image) BeforeDelete(tx *gorm.DB) (err error) {

	if image.Path != "" {
		os.Remove(image.Path)
	}

	return
}
