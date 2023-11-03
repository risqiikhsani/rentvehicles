package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model

	UserID    uint      `json:"user_id" form:"user_id" validate:"required"`
	PostID    uint      `json:"post_id" form:"post_id" validate:"required"`
	StartDate time.Time `json:"start_date" form:"start_date" validate:"required"`
	EndDate   time.Time `json:"end_date" form:"end_date" validate:"required,gtfield=StartDate"`
	// PickupDate    time.Time `json:"pickup_date" form:"pickup_date" validate:"gtefield=StartDate,ltefield=EndDate"`
	// ReturnDate    time.Time `json:"return_date" form:"return_date" validate:"gtefield=PickupDate,ltefield=EndDate"`
	// Status        string  `json:"status" form:"status" gorm:"type:enum('ReadyToPickup', 'Cancelled', 'OnGoing','Done');default:'ReadyToPickup'"`
	PaymentMethod string
	IsCancelled   *bool  `json:"is_cancelled" form:"is_cancelled" gorm:"default:false"`
	CancelReason  string `json:"cancel_reason" form:"cancel_reason"`
	RentDetail    RentDetail
	// Other fields
}

type RentDetail struct {
	gorm.Model
	LicensePlate  string    `json:"license_plate" form:"license_plate"`
	PickupDate    time.Time `json:"pickup_date" form:"pickup_date" validate:"gtefield=StartDate,ltefield=EndDate"`
	ReturnDate    time.Time `json:"return_date" form:"return_date" validate:"gtefield=PickupDate,ltefield=EndDate"`
	DeclineReason string    `json:"decline_reason" form:"decline_reason"`
	Status        string    `json:"status" form:"status" gorm:"default:'ReadyToPickup'"`
	Images        []Image   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Text          string    `json:"text" form:"text"`
	RentID        uint
}

// Declined
// ReadyToPickup
// OnProgress
// Done

func (rent *Rent) BeforeCreate(tx *gorm.DB) (err error) {
	// Fetch the associated Post model by ID
	var post Post
	if err := tx.First(&post, rent.PostID).Error; err != nil {
		return err
	}

	// The error message you're encountering, "invalid operation: operator ! not defined on post.Available (variable of type *bool),"
	// is because the Available field in your Post model is a pointer to a boolean (*bool), and you cannot directly apply the logical NOT (!) operator to a pointer.
	// if !post.Available {
	// 	return fmt.Errorf("Post is not available!")
	// }

	// Check if 'post.Available' is not nil and is set to 'false'
	if post.Available != nil && !*post.Available {
		return fmt.Errorf("Post is not available!")
	}

	return nil
}

func (rent *Rent) AfterCreate(tx *gorm.DB) (err error) {
	// Fetch the associated Post model by ID
	// var post Post
	// if err := tx.First(&post, rent.PostID).Error; err != nil {
	// 	return err
	// }

	// Calculate the updated values
	// remainingUnits := uint(post.Units) - 1
	// post.Units = uint(remainingUnits)
	// post.Available = remainingUnits > 0

	// Update the Post model in the database
	// if err := tx.Save(&post).Error; err != nil {
	// 	return err
	// }

	// Create a RentDetail associated with this Rent
	rentDetail := RentDetail{
		RentID: rent.ID,
	}

	if err := tx.Create(&rentDetail).Error; err != nil {
		return err
	}

	return nil
}
