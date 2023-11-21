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
	userID, _, authenticated := handlers.RequireAuthentication(c, "Admin")
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

	// return error if associated post is not found
	var post models.Post
	if err := models.DB.First(&post, rent.PostID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Associated post not found"})
		return
	}

	// only allow poster (admin) to update rent data
	if post.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this rent detail"})
		return
	}

	// Update the text of the rent
	if err := c.ShouldBind(&existingRentDetail); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	if err := utils.Validate.Struct(existingRentDetail); err != nil {
		errs := utils.TranslateError(err, utils.En)
		c.JSON(400, gin.H{"errors": errs})
		return
	}

	if err := models.DB.Save(&existingRentDetail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Handle image uploads and deletions

	// Parse the multipart form data to handle file uploads
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	imageIDsToDelete := c.PostFormArray("delete_image_ids")

	// Delete selected images from the database and file system
	if len(imageIDsToDelete) > 0 {
		if err := handlers.DeleteRentDetailImages(c, existingRentDetail.ID, imageIDsToDelete); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete images"})
			return
		}
	}

	images := form.File["images"]

	// Handle file uploads and create image records
	if err := handlers.UploadRentDetailImages(c, &existingRentDetail.ID, images); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingRentDetail)
}
