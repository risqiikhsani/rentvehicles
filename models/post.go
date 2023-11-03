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
	Brand         string `json:"brand" form:"brand"  validate:"required"`
	BrandModel    string `json:"brand_model" form:"brand_model"  validate:"required"`
	VehicleType   string `json:"vehicle_type" form:"vehicle_type"  validate:"required"`
	Year          uint   `json:"year" form:"year"  validate:"required,numeric"`
	Transmission  string `json:"transmission" form:"transmission"  validate:"required"`
	FuelType      string `json:"fuel_type" form:"fuel_type"  validate:"required"`
	PricePerDay   uint   `json:"price_per_day" form:"price_per_day"  validate:"required,numeric"`
	PricePerWeek  uint   `json:"price_per_week" form:"price_per_week"  validate:"required,numeric"`
	PricePerMonth uint   `json:"price_per_month" form:"price_per_month"  validate:"required,numeric"`
	Discount      uint   `json:"discount" form:"discount" `
	// Units         uint     `gorm:"default:1" json:"units" form:"units"  validate:"required,numeric"`
	Bookable     *bool   `gorm:"default:true" json:"bookable" form:"bookable" validate:"boolean"`
	BodyColor    string  `json:"body_color" form:"body_color" `
	LicensePlate string  `json:"license_plate" form:"license_plate" `
	Available    *bool   `gorm:"default:true" json:"available" form:"available" validate:"boolean"`
	UserID       uint    `validate:"required,numeric"`
	MainImage    Image   `gorm:"foreignKey:MainPostID"`
	Images       []Image `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // One-to-many relationship with images
	LocationID   uint    `json:"location_id" form:"location_id"  validate:"required,numeric"`
	// Reviews    []Review `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Other fields
}

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
