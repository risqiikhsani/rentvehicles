package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Stars  uint   `json:"stars"`
	Text   string `json:"text"`
	UserID uint
	PostID uint
	RentID uint
}
