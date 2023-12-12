package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/handlers"
	"github.com/risqiikhsani/rentvehicles/models"
	"github.com/risqiikhsani/rentvehicles/utils"
)

func GetLocations(c *gin.Context) {
	var locations []models.Location
	models.DB.Find(&locations)
	c.JSON(200, locations)
}

func GetLocationById(c *gin.Context) {
	locationId := c.Param("location_id")
	var location models.Location
	result := models.DB.First(&location, locationId)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Location not found"})
	}
	c.JSON(200, location)
}

func CreateLocation(c *gin.Context) {
	userID, _, authenticated := handlers.RequireAuthentication(c, "Admin")
	if !authenticated {
		return
	}

	var location models.Location
	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	location.UserID = userID

	if err := utils.Validate.Struct(location); err != nil {
		errs := utils.TranslateError(err, utils.En)
		c.JSON(400, gin.H{"errors": errs})
		return
	}

	// Create the location in the database
	if err := models.DB.Create(&location).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create comment"})
		return
	}
	c.JSON(201, location)

}

func UpdateLocationById(c *gin.Context) {
	// Check if the user is authenticated
	userID, _, authenticated := handlers.RequireAuthentication(c, "Admin")
	if !authenticated {
		return
	}

	locationId := c.Param("location_id")
	var existingLocation models.Location
	if err := models.DB.Where("id = ?", locationId).First(&existingLocation).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "location not found"})
		return
	}

	// Check if the user is the owner of the location
	if existingLocation.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this location"})
		return
	}

	if err := c.ShouldBindJSON(&existingLocation); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := utils.Validate.Struct(existingLocation); err != nil {
		errs := utils.TranslateError(err, utils.En)
		c.JSON(400, gin.H{"errors": errs})
		return
	}

	if err := models.DB.Save(&existingLocation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update location"})
		return
	}

	c.JSON(http.StatusOK, existingLocation)

}

func DeleteLocationById(c *gin.Context) {
	// Check if the user is authenticated
	userID, _, authenticated := handlers.RequireAuthentication(c, "Admin")
	if !authenticated {
		return
	}

	locationId := c.Param("location_id")
	var location models.Location
	result := models.DB.First(&location, locationId)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Comment not found"})
		return
	}

	if location.UserID != userID {
		c.JSON(403, gin.H{"error": "Not authorized to delete this location"})
		return
	}

	models.DB.Delete(&location)

	c.JSON(204, nil)
}
