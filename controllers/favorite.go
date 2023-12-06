package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/handlers"
	"github.com/risqiikhsani/rentvehicles/models"
)

func GetFavorites(c *gin.Context) {
	userID, _, authenticated := handlers.RequireAuthentication(c, "Basic")
	if !authenticated {
		return
	}

	var favorites []models.Favorite
	models.DB.Where("user_id = ?", userID).Find(&favorites)
	c.JSON(200, favorites)
}

func CreateFavorite(c *gin.Context) {
	userID, _, authenticated := handlers.RequireAuthentication(c, "Basic")
	if !authenticated {
		return
	}

	post_id := c.Query("post_id")
	uintVal, err := strconv.ParseUint(post_id, 10, 64)
	if err != nil {
		return
	}

	var favorite models.Favorite

	favorite.UserID = userID
	favorite.PostID = uint(uintVal)

	// Create the favorite in the database
	if err := models.DB.Create(&favorite).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create favorite"})
		return
	}
	c.JSON(201, favorite)

}

func DeleteFavoriteById(c *gin.Context) {
	// Check if the user is authenticated
	userID, _, authenticated := handlers.RequireAuthentication(c, "Basic")
	if !authenticated {
		return
	}

	post_id := c.Query("post_id")
	uintVal, err := strconv.ParseUint(post_id, 10, 64)
	if err != nil {
		return
	}

	var favorite models.Favorite
	if err := models.DB.Where("user_id = ? AND post_id = ?", userID, uintVal).First(&favorite).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Favorite not found"})
		return
	}

	// Delete the favorite by its ID
	if err := models.DB.Delete(&favorite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete favorite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Favorite deleted successfully"})

}
