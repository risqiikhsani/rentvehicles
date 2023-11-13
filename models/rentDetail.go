package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type RentDetail struct {
	gorm.Model
	LicensePlate         string    `json:"license_plate" form:"license_plate"`
	PickupDate           time.Time `json:"pickup_date" form:"pickup_date" validate:"omitempty,gtefield=StartDate,ltefield=EndDate"`
	ReturnDate           time.Time `json:"return_date" form:"return_date" validate:"omitempty,gtefield=PickupDate,ltefield=EndDate"`
	IsPaid               bool      `json:"is_paid" form:"is_paid" gorm:"default:false"`
	EstimatedFinalPrice  uint      `json:"estimated_final_price" form:"estimated_final_price" `
	EstimatedNormalPrice uint      `json:"estimated_normal_price" form:"estimated_normal_price" `
	EstimatedSavedPrice  uint      `json:"estimated_saved_price" form:"estimated_saved_price" `
	RentDays             int       `json:"rent_days" form:"rent_days" `
	DeclineReason        string    `json:"decline_reason" form:"decline_reason"`
	Status               string    `json:"status" form:"status" gorm:"default:'Accepted'"`
	Images               []Image   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Text                 string    `json:"text" form:"text"`
	RentID               uint
}

func (rentDetail *RentDetail) BeforeSave(tx *gorm.DB) (err error) {

	var rent Rent
	if err := tx.First(&rent, rentDetail.RentID).Error; err != nil {
		return err
	}

	if *rent.IsCancelled {
		err = errors.New("can't update, rent was cancelled")
		return err
	}

	return nil
}

func (rentDetail *RentDetail) AfterSave(tx *gorm.DB) (err error) {
	var rent Rent
	if err := tx.First(&rent, rentDetail.RentID).Error; err != nil {
		return err
	}

	var post Post
	if err := tx.First(&post, rent.PostID).Error; err != nil {
		return err
	}

	if rentDetail.Status == "Done" || rentDetail.Status == "Declined" {
		*post.Available = true
	}

	if err := tx.Save(&post).Error; err != nil {
		return err
	}

	return nil
}
