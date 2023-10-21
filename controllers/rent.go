package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/handlers"
	"github.com/risqiikhsani/rentvehicles/models"
)

func GetRents(c *gin.Context) {
	userID, userRole, authenticated := handlers.CheckAuthentication(c)

	if !authenticated {
		return
	}

	var rents []models.Rent

	// if basic user , rents data will be the user's rents history
	if userRole == "basic" {
		models.DB.Preload("Images").Where("user_id = ?", userID).Find(&rents)
		// if admin user, rents data will be the rents data which post_id is admin's
	} else if userRole == "admin" {
		// Assuming an admin can only see rents associated with their own posts.
		subquery := models.DB.Model(&models.Post{}).Select("ID").Where("user_id = ?", userID)
		models.DB.Preload("Images").Where("post_id IN (?)", subquery).Find(&rents)
	} else {
		c.JSON(http.StatusForbidden, gin.H{"message": "Permission denied"})
		return
	}

	c.JSON(http.StatusOK, rents)
}

func GetRentById(c *gin.Context) {
	rent_id := c.Param("rent_id")

	var rent models.Rent

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

	// Fetch the associated Post based on PostID
	var post models.Post
	if err := models.DB.First(&post, rent.PostID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	// Check if the associated Post is available
	if !post.Available {
		c.JSON(400, gin.H{"error": "Post is not available"})
		return
	}

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

	// return error if associated post is not found
	var associatedPost models.Post
	if err := models.DB.First(&associatedPost, existingRent.PostID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Associated post not found"})
		return
	}

	// only allow poster to update rent data
	if associatedPost.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this rent"})
		return
	}

	// Update the text of the rent

	if err := c.ShouldBind(&existingRent); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	if err := models.DB.Save(&existingRent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update rent"})
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
		if err := handlers.DeleteRentImages(c, existingRent.ID, imageIDsToDelete); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete images"})
			return
		}
	}

	files := form.File["files"]

	// Handle file uploads and create image records
	if err := handlers.UploadRentImages(c, &existingRent.ID, files); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingRent)
}

func DeleteRentById(c *gin.Context) {

	// Check if the user is authenticated
	userID, _, authenticated := handlers.CheckAuthentication(c)
	if !authenticated {
		return
	}

	rent_id := c.Param("rent_id")
	var rent models.Rent
	if err := models.DB.First(&rent, rent_id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Rent not found"})
		return
	}

	// if associated post is not found , return error
	var associatedPost models.Post
	if err := models.DB.First(&associatedPost, rent.PostID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Associated post not found"})
		return
	}

	// only customer or poster can delete rent data
	if rent.UserID != userID || associatedPost.UserID != userID {
		c.JSON(403, gin.H{"error": "Not authorized to delete this rent"})
		return
	}

	// delete rent data
	models.DB.Delete(&rent)

	c.JSON(204, nil)
}
