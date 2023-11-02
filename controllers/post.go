package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/handlers"
	"github.com/risqiikhsani/rentvehicles/models"
	"github.com/risqiikhsani/rentvehicles/utils"
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
	if result := models.DB.Preload("Images").Preload("MainImage").First(&post, postId).Error; result != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(200, post)
}

func CreatePost(c *gin.Context) {

	// Check if the user is authenticated
	userID, _, authenticated := handlers.RequireAuthentication(c, "admin")
	if !authenticated {
		return
	}

	// Parse the multipart form data to handle file uploads

	var post models.Post

	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}
	// Set the post's user ID to the authenticated user
	post.UserID = userID

	if err := utils.Validate.Struct(post); err != nil {
		errs := utils.TranslateError(err, utils.En)
		c.JSON(400, gin.H{"errors": errs})
		return
	}

	// Create the post in the database
	if err := models.DB.Create(&post).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create post"})
		return
	}

	// Handle file uploads
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	image := form.File["image"]
	images := form.File["images"]

	// Handle file uploads and create image records
	if err := handlers.UploadMainPostImage(c, &post.ID, image[0]); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Handle file uploads and create image records
	if err := handlers.UploadPostImages(c, &post.ID, images); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, post)
}

func UpdatePostById(c *gin.Context) {
	// Check if the user is authenticated
	userID, _, authenticated := handlers.RequireAuthentication(c, "admin")
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

	// Update the text of the post
	if err := c.ShouldBind(&existingPost); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	if err := utils.Validate.Struct(existingPost); err != nil {
		errs := utils.TranslateError(err, utils.En)
		c.JSON(400, gin.H{"errors": errs})
		return
	}

	if err := models.DB.Save(&existingPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
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
		if err := handlers.DeletePostImages(c, existingPost.ID, imageIDsToDelete); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete images"})
			return
		}
	}

	image := form.File["image"]

	// Handle file uploads and create image records
	if err := handlers.UploadMainPostImage(c, &existingPost.ID, image[0]); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	files := form.File["files"]

	// Handle file uploads and create image records
	if err := handlers.UploadPostImages(c, &existingPost.ID, files); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingPost)
}

func DeletePostById(c *gin.Context) {

	// Check if the user is authenticated
	userID, _, authenticated := handlers.RequireAuthentication(c, "admin")
	if !authenticated {
		return
	}

	postId := c.Param("post_id")
	var post models.Post
	if result := models.DB.First(&post, postId).Error; result != nil {
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
