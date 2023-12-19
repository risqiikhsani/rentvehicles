package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/handlers"
	"github.com/risqiikhsani/rentvehicles/models"
	"github.com/risqiikhsani/rentvehicles/utils"
	"gorm.io/gorm/clause"
)

// Implement other route handlers similarly
func GetPosts(c *gin.Context) {
	var posts []models.Post
	searchQuery := c.Query("search")
	sortBy := c.DefaultQuery("sort", "")
	order := c.DefaultQuery("order", "asc")
	filterBrand := c.Query("filter_brand")
	filterFuelType := c.Query("filter_fuel_type")
	filterTransmission := c.Query("filter_transmission")
	filterVehicleType := c.Query("filter_vehicle_type")

	query := models.DB.Preload(clause.Associations)

	// If a search query is provided, filter the posts based on brand or brand_model
	if searchQuery != "" {
		query = query.Where("LOWER(brand) LIKE ? OR LOWER(brand_model) LIKE ?", "%"+strings.ToLower(searchQuery)+"%", "%"+strings.ToLower(searchQuery)+"%")
	}

	// Apply filters
	if filterBrand != "" {
		query = query.Where("LOWER(brand) = ?", strings.ToLower(filterBrand))
	}
	if filterFuelType != "" {
		query = query.Where("LOWER(fuel_type) = ?", strings.ToLower(filterFuelType))
	}
	if filterTransmission != "" {
		query = query.Where("LOWER(transmission) = ?", strings.ToLower(filterTransmission))
	}
	if filterVehicleType != "" {
		query = query.Where("LOWER(vehicle_type) = ?", strings.ToLower(filterVehicleType))
	}

	// Sort the posts based on the specified field and order
	if sortBy != "" {
		orderClause := sortBy + " " + strings.ToUpper(order)
		query = query.Order(orderClause)
	}

	if err := query.Find(&posts).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No posts found"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func GetMyPosts(c *gin.Context) {
	// Check if the user is authenticated
	userID, _, authenticated := handlers.RequireAuthentication(c, "Admin")
	if !authenticated {
		return
	}

	var posts []models.Post
	if err := models.DB.Where("user_id = ?", userID).Preload(clause.Associations).Find(&posts).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No posts found"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func GetPostById(c *gin.Context) {
	postID := c.Param("post_id")
	var post models.Post

	if err := models.DB.Preload("Images").Preload("MainImage").First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func CreatePost(c *gin.Context) {

	// Check if the user is authenticated
	userID, _, authenticated := handlers.RequireAuthentication(c, "Admin")
	if !authenticated {
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

	mainImage := form.File["main_image"]
	images := form.File["images"]

	// Check if mainImage slice is not empty
	if len(mainImage) == 0 {
		// Handle the case when mainImage is empty, e.g., return an error or take appropriate action.
		c.JSON(400, gin.H{"error": "Main image is missing"})
		return
	}

	// Create the post only if the main image is present
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

	// Handle file uploads and create image records
	if err := handlers.UploadMainPostImage(c, &post.ID, mainImage[0]); err != nil {
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
	userID, _, authenticated := handlers.RequireAuthentication(c, "Admin")
	if !authenticated {
		return
	}

	// Get the post ID from the URL parameters
	postID := c.Param("post_id")

	// Check if the post exists
	var existingPost models.Post
	if err := models.DB.Preload(clause.Associations).Where("id = ?", postID).First(&existingPost).Error; err != nil {
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

	main_image := form.File["main_image"]
	images := form.File["images"]

	// Check if main_image slice is not empty
	if len(main_image) > 0 {
		// Handle file uploads and create image records
		if err := handlers.UploadMainPostImage(c, &existingPost.ID, main_image[0]); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	// Handle file uploads and create image records
	if err := handlers.UploadPostImages(c, &existingPost.ID, images); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingPost)
}

func DeletePostById(c *gin.Context) {

	// Check if the user is authenticated
	userID, _, authenticated := handlers.RequireAuthentication(c, "Admin")
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
