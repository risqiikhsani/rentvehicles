package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/handlers"
	"github.com/risqiikhsani/rentvehicles/models"
	"github.com/risqiikhsani/rentvehicles/utils"
)

func UpdateRentDetailById(c *gin.Context) {
	// Check if the user is authenticated
	userID, _, authenticated := handlers.RequireAuthentication(c, "admin")
	if !authenticated {
		return
	}

	// Get the rent ID from the URL parameters
	rent_detail_id := c.Param("rent_detail_id")

	// Check if the rent exists
	var existingRentDetail models.RentDetail
	if err := models.DB.Where("id = ?", rent_detail_id).First(&existingRentDetail).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rent not found"})
		return
	}

	// return error if associated rent is not found
	var rent models.Rent
	if err := models.DB.First(&rent, existingRentDetail.RentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Associated rent not found"})
		return
	}

	if *rent.IsCancelled {
		c.JSON(http.StatusNotFound, gin.H{"error": "can't update, the rent has been cancelled by consumer."})
		return
	}

	// return error if associated post is not found
	var post models.Post
	if err := models.DB.First(&post, rent.PostID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Associated post not found"})
		return
	}

	// only allow poster (admin) to update rent data
	if rent.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this rent detail"})
		return
	}

	// Update the text of the rent
	if err := c.ShouldBindJSON(&existingRentDetail); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	if err := utils.Validate.Struct(existingRentDetail); err != nil {
		errs := utils.TranslateError(err, utils.En)
		c.JSON(400, gin.H{"errors": errs})
		return
	}

	if err := models.DB.Save(&existingRentDetail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update rent"})
		return
	}

	c.JSON(http.StatusOK, existingRentDetail)
}
