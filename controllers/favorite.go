package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/handlers"
	"github.com/risqiikhsani/rentvehicles/models"
	"gorm.io/gorm/clause"
)

func GetFavorites(c *gin.Context) {
	userID, _, authenticated := handlers.RequireAuthentication(c, "")
	if !authenticated {
		return
	}

	var favorites []models.Favorite

	if err := models.DB.Where("user_id = ?", userID).Find(&favorites).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No posts found"})
		return
	}

	c.JSON(http.StatusOK, favorites)
}

func GetFavoritePosts(c *gin.Context) {
	userID, _, authenticated := handlers.RequireAuthentication(c, "")
	if !authenticated {
		return
	}

	var favorites []models.Favorite
	var posts []models.Post

	// Fetch all favorites for the user
	if err := models.DB.Where("user_id = ?", userID).Find(&favorites).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No posts found"})
		return
	}

	// Extract post IDs from the favorites
	var postIDs []uint
	for _, fav := range favorites {
		postIDs = append(postIDs, fav.PostID)
	}

	// Fetch all posts that have been favorited by the user
	if err := models.DB.Preload(clause.Associations).Where("id IN (?)", postIDs).Find(&posts).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No posts found"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func CreateFavorite(c *gin.Context) {
	userID, _, authenticated := handlers.RequireAuthentication(c, "")
	if !authenticated {
		return
	}

	postID := c.Query("post_id")
	uintVal, err := strconv.ParseUint(postID, 10, 64)
	if err != nil {
		return
	}

	var existingPost models.Post
	if err := models.DB.Where("id = ?", uint(uintVal)).First(&existingPost).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Check if the favorite already exists for the user and post
	var existingFavorite models.Favorite
	if err := models.DB.Where("user_id = ? AND post_id = ?", userID, uint(uintVal)).First(&existingFavorite).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Favorite already exists"})
		return
	}

	favorite := models.Favorite{
		UserID: userID,
		PostID: uint(uintVal),
	}

	// Create the favorite in the database
	if err := models.DB.Create(&favorite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create favorite"})
		return
	}

	c.JSON(http.StatusCreated, favorite)
}
func DeleteFavoriteById(c *gin.Context) {
	// Check if the user is authenticated
	userID, _, authenticated := handlers.RequireAuthentication(c, "")
	if !authenticated {
		return
	}

	post_id := c.Query("post_id")
	uintVal, err := strconv.ParseUint(post_id, 10, 64)
	if err != nil {
		return
	}

	// if post is not found , error
	var existingPost models.Post
	if err := models.DB.Where("id = ?", uint(uintVal)).First(&existingPost).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// if favorite is not found, error
	var favorite models.Favorite
	if err := models.DB.Where("user_id = ? AND post_id = ?", userID, uint(uintVal)).First(&favorite).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Favorite not found"})
		return
	}

	// if user is not the favorite creator, error
	if favorite.UserID != userID {
		c.JSON(403, gin.H{"error": "Not authorized to delete this post"})
		return
	}

	// Delete the favorite by its ID
	if err := models.DB.Delete(&favorite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete favorite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Favorite deleted successfully"})

}
