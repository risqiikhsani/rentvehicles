package handlers

import (
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/models"
)

func DeleteImages(c *gin.Context, postID uint, imageIDs []string) error {
	for _, imageIDToDelete := range imageIDs {
		var imageToDelete models.Image
		if err := models.DB.Where("id = ? AND post_id = ?", imageIDToDelete, postID).First(&imageToDelete).Error; err != nil {
			return err
		}

		// Remove the image file from the file system
		if err := os.Remove(imageToDelete.Path); err != nil {
			return err
		}

		// Delete the image record from the database
		if err := models.DB.Delete(&imageToDelete).Error; err != nil {
			return err
		}
	}

	return nil
}

func UploadImages(c *gin.Context, postID uint, files []*multipart.FileHeader) error {

	for _, fileHeader := range files {
		// Get the file name and path
		filename := filepath.Base(fileHeader.Filename)
		filePath := filepath.Join("static/images", filename)

		// Save the uploaded file to the specified path
		if err := c.SaveUploadedFile(fileHeader, filePath); err != nil {
			return err
		}

		// Assuming you want to store the file paths in the database,
		// you can create an Image model and store the filePath in it.
		// Here's a simplified example:

		image := models.Image{
			Path:   filePath,
			PostID: postID, // Link the image to the post
		}

		if err := models.DB.Create(&image).Error; err != nil {
			return err
		}
	}
	return nil
}
