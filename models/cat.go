package models

import (
	"gorm.io/gorm"
)

type CatDatabase interface {
	GetCats() ([]Cat, error)
	GetCatByID(catID string) (*Cat, error)
	CreateCat(cat *Cat) error
	UpdateCat(cat *Cat) error
	DeleteCat(cat *Cat) error
}

type Cat struct {
	gorm.Model
	Name   string `json:"name" form:"name" binding:"required" validate:"required,min=4,max=15"`
	Age    uint   `json:"age" form:"age" validate:"required,numeric,min=1"`
	Text   string `json:"text" form:"text" validate:"required"`
	UserID uint
	// Images []Image `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // One-to-many relationship with images
	// Other fields
}

func (db *MyDatabase) GetCats() ([]Cat, error) {
	var cats []Cat
	if err := db.Find(&cats).Error; err != nil {
		return nil, err
	}
	return cats, nil
}

func (db *MyDatabase) GetCatByID(catID string) (*Cat, error) {
	var cat Cat
	if err := db.First(&cat, catID).Error; err != nil {
		return nil, err
	}
	return &cat, nil
}

func (db *MyDatabase) CreateCat(cat *Cat) error {
	if err := db.Create(cat).Error; err != nil {
		return err
	}
	return nil
}

func (db *MyDatabase) UpdateCat(cat *Cat) error {
	if err := db.Save(cat).Error; err != nil {
		return err
	}
	return nil
}

func (db *MyDatabase) DeleteCat(cat *Cat) error {
	if err := db.Delete(cat).Error; err != nil {
		return err
	}
	return nil
}
