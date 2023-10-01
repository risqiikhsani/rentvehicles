package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/handlers"
	"github.com/risqiikhsani/rentvehicles/models"
)

// Implement other route handlers similarly
func GetPosts(c *gin.Context) {
	var posts []models.Post
	// Find all posts
	// models.DB.Find(&posts)
	// Find all posts and preload their associated images
	models.DB.Preload("Images").Find(&posts)
	c.JSON(200, posts)
}

func GetPostById(c *gin.Context) {
	postId := c.Param("post_id")

	var post models.Post

	// Find the post by ID
	// result := models.DB.First(&post, postId)
	// Find the post by ID and preload its associated images
	result := models.DB.Preload("Images").First(&post, postId)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(200, post)
}

func UpdatePostById(c *gin.Context) {
	// Check if the user is authenticated
	userID, authenticated := handlers.CheckAuthentication(c)
	if !authenticated {
		return
	}

	// Get the post ID from the URL parameters
	postID := c.Param("post_id")

	// Check if the post exists
	var existingPost models.Post
	if err := models.DB.Preload("Images").Where("id = ?", postID).First(&existingPost).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Check if the user is the owner of the post
	if existingPost.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this post"})
		return
	}

	// Parse the multipart form data to handle file uploads
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Update the text of the post
	existingPost.Text = c.PostForm("text")

	if err := models.DB.Save(&existingPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
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
		if err := handlers.DeleteImages(c, existingPost.ID, imageIDsToDelete); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete images"})
			return
		}
	}

	files := form.File["files"]

	// Handle file uploads and create image records
	if err := handlers.UploadImages(c, existingPost.ID, files); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// // Delete selected images from the database and file system
	// for _, imageIDToDelete := range imageIDsToDelete {
	// 	var imageToDelete models.Image
	// 	if err := models.DB.Where("id = ? AND post_id = ?", imageIDToDelete, existingPost.ID).First(&imageToDelete).Error; err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Image not found or does not belong to the post"})
	// 		return
	// 	}

	// 	// Remove the image file from the file system
	// 	if err := os.Remove(imageToDelete.Path); err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete image file"})
	// 		return
	// 	}

	// 	// Delete the image record from the database
	// 	if err := models.DB.Delete(&imageToDelete).Error; err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete image record"})
	// 		return
	// 	}
	// }

	// files := form.File["files"]

	// for _, fileHeader := range files {
	// 	// Get the file name and path
	// 	filename := filepath.Base(fileHeader.Filename)
	// 	filePath := filepath.Join("static/images", filename)

	// 	// Save the uploaded file to the specified path
	// 	if err := c.SaveUploadedFile(fileHeader, filePath); err != nil {
	// 		c.JSON(400, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	// Assuming you want to store the file paths in the database,
	// 	// you can create an Image model and store the filePath in it.
	// 	// Here's a simplified example:

	// 	image := models.Image{
	// 		Path:   filePath,
	// 		PostID: existingPost.ID, // Link the image to the post
	// 	}

	// 	if err := models.DB.Create(&image).Error; err != nil {
	// 		c.JSON(500, gin.H{"error": "Failed to create image record"})
	// 		return
	// 	}
	// }

	c.JSON(http.StatusOK, existingPost)
}

func CreatePost(c *gin.Context) {

	// Check if the user is authenticated
	userID, authenticated := handlers.CheckAuthentication(c)
	if !authenticated {
		return
	}

	// Parse the multipart form data to handle file uploads
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var post models.Post
	// Retrieve the text field from the form data
	post.Text = c.PostForm("text")

	// Set the post's user ID to the authenticated user
	post.UserID = userID

	// Create the post in the database
	if err := models.DB.Create(&post).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create post"})
		return
	}

	// Handle file uploads

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	files := form.File["files"]

	// Handle file uploads and create image records
	if err := handlers.UploadImages(c, post.ID, files); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, post)
}

func DeletePostById(c *gin.Context) {

	// Check if the user is authenticated
	userID, authenticated := handlers.CheckAuthentication(c)
	if !authenticated {
		return
	}

	postId := c.Param("post_id")
	var post models.Post
	result := models.DB.First(&post, postId)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	if post.UserID != userID {
		c.JSON(403, gin.H{"error": "Not authorized to delete this post"})
		return
	}

	models.DB.Delete(&post)

	c.JSON(204, nil)
}
