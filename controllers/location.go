package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/handlers"
	"github.com/risqiikhsani/rentvehicles/models"
)

func getLocations(c *gin.Context) {
	var locations []models.GoogleMapLocation
	models.DB.Find(&locations)
	c.JSON(200, locations)
}

func getLocationById(c *gin.Context) {
	locationId := c.Param("location_id")
	var location models.GoogleMapLocation
	result := models.DB.First(&location, locationId)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Location not found"})
	}
	c.JSON(200, location)
}

func CreateLocation(c *gin.Context) {
	userID, userRole, authenticated := handlers.CheckAuthentication(c)
	if !authenticated {
		return
	}

	if userRole != "admin" {
		return
	}

	var location models.GoogleMapLocation
	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	location.UserID = userID
	// Create the location in the database
	if err := models.DB.Create(&location).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create comment"})
		return
	}
	c.JSON(201, location)

}

func UpdateLocationById(c *gin.Context) {
	// Check if the user is authenticated
	userID, _, authenticated := handlers.CheckAuthentication(c)
	if !authenticated {
		return
	}

	locationId := c.Param("location_id")
	var existingLocation models.GoogleMapLocation
	if err := models.DB.Where("id = ?", locationId).First(&existingLocation).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Check if the user is the owner of the comment
	if existingLocation.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this comment"})
		return
	}

	if err := c.ShouldBindJSON(&existingLocation); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Save(&existingLocation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	c.JSON(http.StatusOK, existingLocation)

}

func DeleteLocationById(c *gin.Context) {
	// Check if the user is authenticated
	userID, _, authenticated := handlers.CheckAuthentication(c)
	if !authenticated {
		return
	}

	locationId := c.Param("location_id")
	var location models.GoogleMapLocation
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