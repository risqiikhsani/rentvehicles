package models

import (
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
	Text         string `json:"text" binding:"required"`
	UserID       uint   // default colum name will be user_id, you can specify it with `gorm:"column:desiredname"`
	PostID       uint
	StartDate    time.Time
	EndDate      time.Time
	PickupDate   time.Time
	ReturnDate   time.Time
	LicensePlate string `json:"license_plate"`
	Status       string `gorm:"default:'ReadyToPickup'" json:"status"`
	Images       []Image
	// Other fields
}

// Cancelled
// ReadyToPickup
// OnGoing
// Done
