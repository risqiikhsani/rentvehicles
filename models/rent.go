package models

import (
	"errors"
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
	PaymentMethod string `json:"payment_method" form:"payment_method" gorm:"default:'Paylater'" validate:"omitempty,oneof=Paylater PayInFront"`
	IsCancelled   *bool  `json:"is_cancelled" form:"is_cancelled" gorm:"default:false"`
	CancelReason  string `json:"cancel_reason" form:"cancel_reason"`
	DiscountCode  string `json:"discount_voucher" form:"discount_voucher"`
	Readonly      *bool  `json:"readonly" form:"readonly" gorm:"default:false"`
	RentDetail    RentDetail
	// Other fields
}

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
	finalPrice, normalPrice, savedPrice, _ := post.CalculateRentalPrice(uint(totalDays))
	// Create a RentDetail associated with this Rent
	rentDetail := RentDetail{
		PickupDate:           nil,
		ReturnDate:           nil,
		RentDays:             totalDays,
		RentID:               rent.ID,
		EstimatedFinalPrice:  finalPrice,
		EstimatedNormalPrice: normalPrice,
		EstimatedSavedPrice:  savedPrice,
	}

	if err := tx.Create(&rentDetail).Error; err != nil {
		return err
	}

	return nil
}

func (rent *Rent) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("Rent's before update is running")

	// this will make error "rent is read only" when RentDetail's AfterUpdate hook was called
	// eventhough the rent's readonly was false in the database.
	// if rent.Readonly {
	// 	err = errors.New("Rent is read only")
	// 	return err
	// }

	// to solve the problem above, we have to fetch the data from db first, then check it
	// Retrieve the existing record from the database to compare with the updated values
	var existingRent Rent
	if err := tx.First(&existingRent, rent.ID).Error; err != nil {
		return err
	}

	// Check if the rent is marked as read-only
	if *existingRent.Readonly {
		fmt.Println("if rent readonly was true, can't update !")
		err = errors.New("Rent is read only")
		return err
	}

	// If the rent is cancelled, mark it as read-only
	if *rent.IsCancelled {
		fmt.Println("if rent is cancelled , readonly = true")
		*rent.Readonly = true
	}

	return nil
}

// AfterUpdate is a method that is called after a Rent model is updated
// It updates the associated Post model based on the updated Rent status
func (rent *Rent) AfterUpdate(tx *gorm.DB) (err error) {
	fmt.Println("Rent's after update is running")
	// Fetch the associated Post model by ID
	var post Post
	if err := tx.First(&post, rent.PostID).Error; err != nil {
		return err
	}

	// var existingRent Rent
	// if err := tx.First(&existingRent, rent.ID).Error; err != nil {
	// 	return err
	// }

	// If the Rent is cancelled, set the Post as available
	if *rent.IsCancelled {
		fmt.Println("If rent is cancelled, post available = true")
		// *rent.Readonly = true
		*post.Available = true
	}

	// Save the updated Post model
	if err := tx.Save(&post).Error; err != nil {
		return err
	}

	// // Save the updated Post model
	// if err := tx.Save(&existingRent).Error; err != nil {
	// 	return err
	// }

	return
}
