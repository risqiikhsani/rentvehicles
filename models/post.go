package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Brand                      string `json:"brand" form:"brand"  validate:"required"`
	BrandModel                 string `json:"brand_model" form:"brand_model"  validate:"required"`
	VehicleType                string `json:"vehicle_type" form:"vehicle_type"  validate:"required"`
	Year                       uint   `json:"year" form:"year"  validate:"required,numeric"`
	Transmission               string `json:"transmission" form:"transmission"  validate:"required"`
	FuelType                   string `json:"fuel_type" form:"fuel_type"  validate:"required"`
	PricePerDay                uint   `json:"price_per_day" form:"price_per_day"  validate:"required,numeric"`
	PricePerWeek               uint   `json:"price_per_week" form:"price_per_week"  validate:"required,numeric"`
	PricePerMonth              uint   `json:"price_per_month" form:"price_per_month"  validate:"required,numeric"`
	PricePerDayAfterDiscount   uint   `json:"price_per_day_after_discount" form:"price_per_day_after_discount" `
	PricePerWeekAfterDiscount  uint   `json:"price_per_week_after_discount" form:"price_per_week_after_discount" `
	PricePerMonthAfterDiscount uint   `json:"price_per_month_after_discount" form:"price_per_month_after_discount" `
	DiscountPercentage         uint   `json:"discount_percentage" form:"discount_percentage" `
	// Units         uint     `gorm:"default:1" json:"units" form:"units"  validate:"required,numeric"`
	Bookable     *bool   `gorm:"default:true" json:"bookable" form:"bookable"`
	BodyColor    string  `json:"body_color" form:"body_color" `
	LicensePlate string  `json:"license_plate" form:"license_plate" `
	Available    *bool   `gorm:"default:true" json:"available" form:"available"`
	UserID       uint    `validate:"required,numeric"`
	MainImage    Image   `gorm:"foreignKey:MainPostID"`
	Images       []Image `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // One-to-many relationship with images
	LocationID   uint    `json:"location_id" form:"location_id"  validate:"required,numeric"`
	// Reviews    []Review `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Other fields
}

func (post *Post) CalculateRentalPrice(rentDays uint) (uint, uint, uint, error) {
	if rentDays == 0 {
		return 0, 0, 0, nil
	}

	// Calculate the total price based on rental duration
	var totalPrice uint
	var discountedPrice uint
	savedPrice := 0

	if rentDays >= 30 {
		// Calculate the number of months and remaining days
		months := rentDays / 30
		remainingDays := rentDays % 30

		// Calculate the total price
		totalMonthlyPrice := months * post.PricePerMonth
		totalDailyPrice := remainingDays * post.PricePerDay

		totalPrice = totalMonthlyPrice + totalDailyPrice
	} else if rentDays >= 7 {
		// Calculate the number of weeks and remaining days
		weeks := rentDays / 7
		remainingDays := rentDays % 7

		// Calculate the total price
		totalWeeklyPrice := weeks * post.PricePerWeek
		totalDailyPrice := remainingDays * post.PricePerDay

		totalPrice = totalWeeklyPrice + totalDailyPrice
	} else {
		// If the rental duration is less than a week, charge the daily rate
		totalPrice = rentDays * post.PricePerDay
	}

	if post.DiscountPercentage != 0 {
		discountedPrice = totalPrice - (totalPrice * post.DiscountPercentage / 100)
		savedPrice = int(totalPrice * post.DiscountPercentage / 100)
		return discountedPrice, totalPrice, uint(savedPrice), nil
	}

	return totalPrice, totalPrice, uint(savedPrice), nil
}

func (post *Post) BeforeSave(tx *gorm.DB) (err error) {

	post.PricePerDayAfterDiscount = post.PricePerDay
	post.PricePerWeekAfterDiscount = post.PricePerWeek
	post.PricePerMonthAfterDiscount = post.PricePerMonth

	if post.DiscountPercentage != 0 {
		post.PricePerDayAfterDiscount = post.PricePerDay - (post.PricePerDay * post.DiscountPercentage / 100)
		post.PricePerWeekAfterDiscount = post.PricePerWeek - (post.PricePerWeek * post.DiscountPercentage / 100)
		post.PricePerMonthAfterDiscount = post.PricePerMonth - (post.PricePerMonth * post.DiscountPercentage / 100)
	}

	return nil
}

func (post *Post) AfterDelete(tx *gorm.DB) (err error) {
	// First, fetch all associated images
	var images []Image
	if err := tx.Model(post).Association("Images").Find(&images); err != nil {
		return err
	}

	// Delete each associated image
	for _, image := range images {
		if err := tx.Unscoped().Delete(&image).Error; err != nil {
			return err
		}
	}

	return nil
}

func (post *Post) BeforeDelete(tx *gorm.DB) (err error) {

	// if there is a rent for this post , that is on going . admin can't delete or update

	return nil
}

func (post *Post) BeforeUpdate(tx *gorm.DB) (err error) {

	// if there is a rent for this post , that is on going . admin can't delete or update

	return nil
}
