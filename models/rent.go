package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	Text         string    `json:"text" form:"text"`
	UserID       uint      `json:"user_id" form:"user_id" validate:"required"`
	PostID       uint      `json:"post_id" form:"post_id" validate:"required"`
	StartDate    time.Time `json:"start_date" form:"start_date" validate:"required"`
	EndDate      time.Time `json:"end_date" form:"end_date" validate:"required,gtfield=StartDate"`
	PickupDate   time.Time `json:"pickup_date" form:"pickup_date" validate:"gtefield=StartDate,ltefield=EndDate"`
	ReturnDate   time.Time `json:"return_date" form:"return_date" validate:"gtefield=PickupDate,ltefield=EndDate"`
	LicensePlate string    `json:"license_plate" form:"license_plate"`
	Status       string    `json:"status" form:"status" gorm:"default:'ReadyToPickup'"`
	Images       []Image   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Other fields
}

// Cancelled
// ReadyToPickup
// OnGoing
// Done

func (rent *Rent) BeforeCreate(tx *gorm.DB) (err error) {
	// Fetch the associated Post model by ID
	var post Post
	if err := tx.First(&post, rent.PostID).Error; err != nil {
		return err
	}
	// If post.Units is zero or post.Available is false, cancel the creation of rent
	if post.Units < 1 || !post.Available {
		return fmt.Errorf("cannot create Rent, item is not available")
	}

	return nil
}

func (rent *Rent) AfterCreate(tx *gorm.DB) (err error) {
	// Fetch the associated Post model by ID
	var post Post
	if err := tx.First(&post, rent.PostID).Error; err != nil {
		return err
	}

	// Calculate the updated values
	remainingUnits := uint(post.Units) - 1
	post.Units = uint(remainingUnits)
	post.Available = remainingUnits > 0

	// Update the Post model in the database
	if err := tx.Save(&post).Error; err != nil {
		return err
	}

	return nil
}
