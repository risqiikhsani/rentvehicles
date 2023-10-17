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
	result := models.DB.Preload("Images").Preload("Reviews").First(&post, postId)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(200, post)
}

func CreatePost(c *gin.Context) {

	// Check if the user is authenticated
	userID, userRole, authenticated := handlers.CheckAuthentication(c)
	if !authenticated {
		return
	}

	if userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to create post"})
		return
	}

	// Parse the multipart form data to handle file uploads
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var post models.Post
	// Retrieve the text field from the form data)
	post.Brand = c.PostForm("brand")
	post.BrandModel = c.PostForm("brand_model")
	post.FuelType = c.PostForm("fuel_type")
	post.LocationID, err = utils.ConvertToUint(c.PostForm("location_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	post.Transmission = c.PostForm("transmission")
	post.Units, err = utils.ConvertToUint(c.PostForm("units"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	post.VehicleType = c.PostForm("vehicle_type")
	post.Year, err = utils.ConvertToUint(c.PostForm("year"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	post.PricePerDay, err = utils.ConvertToUint(c.PostForm("price_per_day"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	post.PricePerWeek, err = utils.ConvertToUint(c.PostForm("price_per_week"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	post.PricePerMonth, err = utils.ConvertToUint(c.PostForm("price_per_month"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	post.Discount, err = utils.ConvertToUint(c.PostForm("discount"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

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

func UpdatePostById(c *gin.Context) {
	// Check if the user is authenticated
	userID, _, authenticated := handlers.CheckAuthentication(c)
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
	existingPost.Brand = c.PostForm("brand")
	existingPost.BrandModel = c.PostForm("brand_model")
	existingPost.FuelType = c.PostForm("fuel_type")
	existingPost.LocationID, err = utils.ConvertToUint(c.PostForm("location_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	existingPost.Transmission = c.PostForm("transmission")
	existingPost.Units, err = utils.ConvertToUint(c.PostForm("units"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	existingPost.VehicleType = c.PostForm("vehicle_type")
	existingPost.Year, err = utils.ConvertToUint(c.PostForm("year"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	existingPost.PricePerDay, err = utils.ConvertToUint(c.PostForm("price_per_day"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	existingPost.PricePerWeek, err = utils.ConvertToUint(c.PostForm("price_per_week"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	existingPost.PricePerMonth, err = utils.ConvertToUint(c.PostForm("price_per_month"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	existingPost.Discount, err = utils.ConvertToUint(c.PostForm("discount"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

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

	c.JSON(http.StatusOK, existingPost)
}

func DeletePostById(c *gin.Context) {

	// Check if the user is authenticated
	userID, _, authenticated := handlers.CheckAuthentication(c)
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
