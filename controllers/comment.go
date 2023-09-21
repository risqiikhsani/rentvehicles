package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/contactgo/handlers"
	"github.com/risqiikhsani/contactgo/models"
)

// Implement other route handlers similarly
func GetComments(c *gin.Context) {
	postId := c.Param("post_id")
	var comments []models.Comment
	// Find all comments
	// models.DB.Find(&comments)
	// Find all comments and preload their associated images
	models.DB.Where("post_id = ?", postId).Find(&comments)
	c.JSON(200, comments)
}

func GetCommentById(c *gin.Context) {
	commentId := c.Param("comment_id")

	var comment models.Comment

	// Find the comment by ID
	// result := models.DB.First(&comment, commentId)
	// Find the comment by ID and preload its associated images
	result := models.DB.First(&comment, commentId)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(200, comment)
}

func UpdateCommentById(c *gin.Context) {
	// Check if the user is authenticated
	userID, authenticated := handlers.CheckAuthentication(c)
	if !authenticated {
		return
	}

	// Get the comment ID from the URL parameters
	commentId := c.Param("comment_id")

	// Check if the comment exists
	var existingComment models.Comment
	if err := models.DB.Where("id = ?", commentId).First(&existingComment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Check if the user is the owner of the comment
	if existingComment.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this comment"})
		return
	}

	if err := c.ShouldBindJSON(&existingComment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Save(&existingComment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	c.JSON(http.StatusOK, existingComment)
}

func CreateComment(c *gin.Context) {
	postId := c.Param("post_id")
	postIdUint, err := strconv.ParseUint(postId, 10, 64)

	if err != nil {
		// Handle the error if the conversion fails
		fmt.Println("Error:", err)
		return
	}

	// Check if the user is authenticated
	userID, authenticated := handlers.CheckAuthentication(c)
	if !authenticated {
		return
	}

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	comment.PostID = uint(postIdUint)
	comment.UserID = userID
	// Create the post in the database
	if err := models.DB.Create(&comment).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(201, comment)
}

func DeleteCommentById(c *gin.Context) {

	// Check if the user is authenticated
	userID, authenticated := handlers.CheckAuthentication(c)
	if !authenticated {
		return
	}

	commentId := c.Param("comment_id")
	var comment models.Comment
	result := models.DB.First(&comment, commentId)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Comment not found"})
		return
	}

	if comment.UserID != userID {
		c.JSON(403, gin.H{"error": "Not authorized to delete this comment"})
		return
	}

	models.DB.Delete(&comment)

	c.JSON(204, nil)
}
