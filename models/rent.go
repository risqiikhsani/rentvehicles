package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	Text         string `json:"text"`
	UserID       uint   // default colum name will be user_id, you can specify it with `gorm:"column:desiredname"`
	PostID       uint
	StartDate    time.Time
	EndDate      time.Time
	PickupDate   time.Time
	ReturnDate   time.Time
	LicensePlate string  `json:"license_plate"`
	Status       string  `gorm:"default:'ReadyToPickup'" json:"status"`
	Images       []Image `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Other fields
}

// Cancelled
// ReadyToPickup
// OnGoing
// Done

func (rent *Rent) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("after create hook is running")
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
