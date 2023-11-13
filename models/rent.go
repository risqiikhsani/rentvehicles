package models

import (
	"errors"
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
	PaymentMethod string `json:"payment_method" form:"payment_method" gorm:"default:'paylater'"`
	IsCancelled   *bool  `json:"is_cancelled" form:"is_cancelled" gorm:"default:false"`
	CancelReason  string `json:"cancel_reason" form:"cancel_reason"`
	DiscountCode  string `json:"discount_voucher" form:"discount_voucher"`
	Readonly      bool   `json:"readonly" form:"readonly" gorm:"default:false"`
	RentDetail    RentDetail
	// Other fields
}

// payment method
// paylater
// pay in front

// Declined
// Accepted
// ReadyToPickup
// OnProgress
// Done

func (rent *Rent) CalculateTotalDays() int {
	// Calculate the duration between the StartDate and EndDate
	duration := rent.EndDate.Sub(rent.StartDate)

	// Extract the total number of days from the duration
	totalDays := int(duration.Hours() / 24)

	return totalDays
}

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
		err = errors.New("Post is not available")
		return err
	}

	return nil
}

func (rent *Rent) AfterCreate(tx *gorm.DB) (err error) {
	// Fetch the associated Post model by ID
	var post Post
	if err := tx.First(&post, rent.PostID).Error; err != nil {
		return err
	}

	*post.Available = false

	if err := tx.Save(&post).Error; err != nil {
		return err
	}

	totalDays := rent.CalculateTotalDays()
	estimatedPrice, savedPrice, _ := post.CalculateRentalPrice(uint(totalDays))
	// Create a RentDetail associated with this Rent
	rentDetail := RentDetail{
		RentID:              rent.ID,
		EstimatedPrice:      estimatedPrice,
		EstimatedSavedPrice: savedPrice,
	}

	if err := tx.Create(&rentDetail).Error; err != nil {
		return err
	}

	return nil
}

func (rent *Rent) BeforeUpdate(tx *gorm.DB) (err error) {

	return
}

func (rent *Rent) AfterUpdate(tx *gorm.DB) (err error) {
	// Fetch the associated Post model by ID
	var post Post
	if err := tx.First(&post, rent.PostID).Error; err != nil {
		return err
	}

	if *rent.IsCancelled {
		*post.Available = true
	}

	if err := tx.Save(&post).Error; err != nil {
		return err
	}

	return
}
