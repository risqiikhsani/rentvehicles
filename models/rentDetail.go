package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type RentDetail struct {
	gorm.Model
	PickupDate           *time.Time `json:"pickup_date" form:"pickup_date" validate:"omitempty,gtefield=StartDate,ltefield=EndDate"`
	ReturnDate           *time.Time `json:"return_date" form:"return_date" validate:"omitempty,gtefield=PickupDate,ltefield=EndDate"`
	IsPaid               bool       `json:"is_paid" form:"is_paid" gorm:"default:false"`
	EstimatedFinalPrice  uint       `json:"estimated_final_price" form:"estimated_final_price" `
	EstimatedNormalPrice uint       `json:"estimated_normal_price" form:"estimated_normal_price" `
	EstimatedSavedPrice  uint       `json:"estimated_saved_price" form:"estimated_saved_price" `
	RentDays             int        `json:"rent_days" form:"rent_days" `
	DeclineReason        string     `json:"decline_reason" form:"decline_reason"`
	Status               string     `json:"status" form:"status" gorm:"default:'Accepted'" validate:"omitempty,oneof=Accepted Declined ReadyToPickup OnProgress Done"`
	Images               []Image    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Text                 string     `json:"text" form:"text"`
	RentID               uint
}

// Declined
// Accepted
// ReadyToPickup
// OnProgress
// Done

func (rentDetail *RentDetail) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("RentDetail's before update is running")

	// Additional checks after status validation
	var rent Rent
	if err := tx.First(&rent, rentDetail.RentID).Error; err != nil {
		return err
	}

	if *rent.IsCancelled {
		return errors.New("can't update, rent was cancelled")
	}

	if rentDetail.Status == "OnProgress" || rentDetail.Status == "Done" {
		if !rentDetail.IsPaid {
			return errors.New("can't update, rent is not paid")
		}
	}

	return nil // All checks passed, return nil for no errors
}

func (rentDetail *RentDetail) AfterUpdate(tx *gorm.DB) (err error) {
	fmt.Println("RentDetail's after update is running")
	var existingRent Rent
	if err := tx.First(&existingRent, rentDetail.RentID).Error; err != nil {
		return err
	}

	var existingPost Post
	if err := tx.First(&existingPost, existingRent.PostID).Error; err != nil {
		return err
	}

	if rentDetail.Status == "OnProgress" || rentDetail.Status == "Done" || rentDetail.Status == "Declined" {
		*existingRent.Readonly = true // Assuming Readonly is not a pointer
	}

	if rentDetail.Status == "Done" || rentDetail.Status == "Declined" {
		*existingPost.Available = true // Assuming Available is not a pointer
	}

	if err := tx.Save(&existingPost).Error; err != nil {
		return err
	}

	if err := tx.Save(&existingRent).Error; err != nil {
		return err
	}

	return nil
}

// func (rentDetail *RentDetail) AfterUpdate(tx *gorm.DB) (err error) {
// 	// Fetch the associated Rent model by ID
// 	var rent Rent
// 	if err := tx.First(&rent, rentDetail.RentID).Error; err != nil {
// 		return err
// 	}

// 	// Update Rent's Readonly field to true
// 	rent.Readonly = true

// 	// Save the updated Rent model
// 	if err := tx.Save(&rent).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }
