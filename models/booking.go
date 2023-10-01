package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UserID    uint
	PostID    uint
	StartDate time.Time
	EndDate   time.Time
}
