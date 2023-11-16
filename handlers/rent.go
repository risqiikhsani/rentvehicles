package handlers

import (
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/models"
	"github.com/risqiikhsani/rentvehicles/utils"
)

func UploadRentDetailImages(c *gin.Context, rentDetailID *uint, files []*multipart.FileHeader) error {
	// Check if the rentDetail with the provided rentDetailID exists
	var rentDetail models.RentDetail
	if err := models.DB.First(&rentDetail, rentDetailID).Error; err != nil {
		return err
	}

	for _, fileHeader := range files {
		// Get the file name and path
		// filename := filepath.Base(fileHeader.Filename)

		// Generate a random string for the file name
		randomString := utils.RandomStringUuid()

		// Get the file extension from the original file name
		fileExt := filepath.Ext(fileHeader.Filename)

		// Create a unique file name by combining the random string and the file extension
		filename := randomString + fileExt
		filePath := filepath.Join("static/images", filename)

		// Save the uploaded file to the specified path
		if err := c.SaveUploadedFile(fileHeader, filePath); err != nil {
			return err
		}

		// Assuming you want to store the file paths in the database,
		// you can create an Image model and store the filePath in it.
		// Here's a simplified example:

		image := models.Image{
			Path:         filePath,
			RentDetailID: rentDetailID, // Link the image to the rentDetail
		}

		if err := models.DB.Create(&image).Error; err != nil {
			return err
		}
	}
	return nil
}

func DeleteRentDetailImages(c *gin.Context, rentDetailID uint, imageIDs []string) error {
	for _, imageIDToDelete := range imageIDs {
		var imageToDelete models.Image
		if err := models.DB.Where("id = ? AND rent_detail_id = ?", imageIDToDelete, rentDetailID).First(&imageToDelete).Error; err != nil {
			return err
		}

		// Remove the image file from the file system
		// handled by hooks in model

		// Delete the image record from the database
		if err := models.DB.Delete(&imageToDelete).Error; err != nil {
			return err
		}
	}

	return nil
}
