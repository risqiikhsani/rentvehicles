package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/risqiikhsani/rentvehicles/handlers"
	"github.com/risqiikhsani/rentvehicles/models"
	"github.com/risqiikhsani/rentvehicles/utils"
)

func GetCats(c *gin.Context) {
	var cats []models.Cat
	models.DB.Find(&cats)
	c.JSON(200, cats)
}

func GetCatById(c *gin.Context) {
	catId := c.Param("cat_id")
	var cat models.Cat
	result := models.DB.First(&cat, catId)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Location not found"})
	}
	c.JSON(200, cat)
}

func CreateCat(c *gin.Context) {
	userID, userRole, authenticated := handlers.CheckAuthentication(c)

	if !authenticated {
		return
	}

	if userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to create cat"})
		return
	}

	var cat models.Cat
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := utils.Validate.Struct(cat); err != nil {
		errs := utils.TranslateError(err, utils.En)
		c.JSON(400, gin.H{"errors": errs})
		return
	}

	cat.UserID = userID
	// Create the cat in the database
	if err := models.DB.Create(&cat).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create comment"})
		return
	}
	c.JSON(201, cat)

}

func UpdateCatById(c *gin.Context) {
	// Check if the user is authenticated
	userID, _, authenticated := handlers.CheckAuthentication(c)
	if !authenticated {
		return
	}

	catId := c.Param("cat_id")
	var existingCat models.Cat
	if err := models.DB.Where("id = ?", catId).First(&existingCat).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cat not found"})
		return
	}

	// Check if the user is the owner of the cat
	if existingCat.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this cat"})
		return
	}

	if err := c.ShouldBindJSON(&existingCat); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := utils.Validate.Struct(existingCat); err != nil {
		var errs []string // Change the type to a string slice

		for _, err := range err.(validator.ValidationErrors) {
			// You can create a more descriptive error message if needed
			fieldError := fmt.Sprintf("Field: %s, Error: %s", err.Field(), err.Tag())
			errs = append(errs, fieldError)
		}

		c.JSON(400, gin.H{"errors": errs})
		return
	}

	if err := models.DB.Save(&existingCat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cat"})
		return
	}

	c.JSON(http.StatusOK, existingCat)

}

func DeleteCatById(c *gin.Context) {
	// Check if the user is authenticated
	userID, _, authenticated := handlers.CheckAuthentication(c)
	if !authenticated {
		return
	}

	catId := c.Param("cat_id")
	var cat models.Cat
	result := models.DB.First(&cat, catId)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Comment not found"})
		return
	}

	if cat.UserID != userID {
		c.JSON(403, gin.H{"error": "Not authorized to delete this cat"})
		return
	}

	models.DB.Delete(&cat)

	c.JSON(204, nil)
}
