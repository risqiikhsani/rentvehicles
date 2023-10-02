package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/handlers"
	"github.com/risqiikhsani/rentvehicles/models"
)

// Implement other route handlers similarly
func GetRents(c *gin.Context) {
	var rents []models.Rent
	// Find all rents
	// models.DB.Find(&rents)
	// Find all rents and preload their associated images
	models.DB.Preload("Images").Find(&rents)

	c.JSON(200, rents)
}

func GetRentById(c *gin.Context) {
	rent_id := c.Param("rent_id")

	var rent models.Rent

	// Find the rent by ID
	// result := models.DB.First(&rent, rent_id)
	// Find the rent by ID and preload its associated images
	result := models.DB.Preload("Images").First(&rent, rent_id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Rent not found"})
		return
	}

	c.JSON(200, rent)
}

func CreateRent(c *gin.Context) {
	userID, _, authenticated := handlers.CheckAuthentication(c)
	if !authenticated {
		return
	}

	var rent models.Rent
	if err := c.ShouldBindJSON(&rent); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	rent.UserID = userID
	// Create the rent in the database
	if err := models.DB.Create(&rent).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create comment"})
		return
	}
	c.JSON(201, rent)

}

func UpdateRentById(c *gin.Context) {
	// Check if the user is authenticated
	userID, userRole, authenticated := handlers.CheckAuthentication(c)
	if !authenticated {
		return
	}

	if userRole != "admin" {
		return
	}

	// Get the rent ID from the URL parameters
	rent_id := c.Param("rent_id")

	// Check if the rent exists
	var existingRent models.Rent
	if err := models.DB.Preload("Images").Where("id = ?", rent_id).First(&existingRent).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rent not found"})
		return
	}

	// only allow post creator to update
	var associatedPost models.Post
	if err := models.DB.First(&associatedPost, existingRent.PostID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Associated post not found"})
		return
	}
	if associatedPost.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this rent"})
		return
	}

	// Parse the multipart form data to handle file uploads
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Update the text of the rent
	existingRent.Text = c.PostForm("text")
	pickupDateStr := c.PostForm("pickup_date")
	pickupDate, err := time.Parse("2006-01-02", pickupDateStr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		existingRent.PickupDate = pickupDate
	}
	returnDateStr := c.PostForm("return_date")
	returnDate, err := time.Parse("2006-01-02", returnDateStr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		existingRent.PickupDate = returnDate
	}
	existingRent.LicensePlate = c.PostForm("license_plate")
	existingRent.Status = c.PostForm("status")

	if err := models.DB.Save(&existingRent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update rent"})
		return
	}

	// Handle image uploads and deletions
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	imageIDsToDelete := c.PostFormArray("delete_image_ids")

	// Delete selected images from the database and file system
	if len(imageIDsToDelete) > 0 {
		if err := handlers.DeleteImages(c, existingRent.ID, imageIDsToDelete); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete images"})
			return
		}
	}

	files := form.File["files"]

	// Handle file uploads and create image records
	if err := handlers.UploadImages(c, existingRent.ID, files); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingRent)
}

func DeleteRenttById(c *gin.Context) {

	// Check if the user is authenticated
	userID, _, authenticated := handlers.CheckAuthentication(c)
	if !authenticated {
		return
	}

	rent_id := c.Param("rent_id")
	var rent models.Rent
	result := models.DB.First(&rent, rent_id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Rent not found"})
		return
	}

	// only allow post creator or the rent creator to delete
	var associatedPost models.Post
	if err := models.DB.First(&associatedPost, rent.PostID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Associated post not found"})
		return
	}

	if rent.UserID != userID || associatedPost.UserID != userID {
		c.JSON(403, gin.H{"error": "Not authorized to delete this rent"})
		return
	}

	models.DB.Delete(&rent)

	c.JSON(204, nil)
}
