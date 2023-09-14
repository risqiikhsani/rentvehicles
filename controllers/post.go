package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/contactgo/models"
)

// Implement other route handlers similarly
func GetPosts(c *gin.Context) {
	var posts []models.Post
	models.DB.Find(&posts)
	c.JSON(200, posts)
}

func GetPostById(c *gin.Context) {
	postId := c.Param("id")

	var post models.Post

	// Find the post by ID
	result := models.DB.First(&post, postId)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(200, post)
}

func UpdatePostById(c *gin.Context) {
	var post models.Post
	postId := c.Param("id")
	userIDValue, exists := c.Get("userID")

	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	userId := userIDValue.(uint)

	result := models.DB.First(&post, postId)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	if post.UserID != userId {
		c.JSON(403, gin.H{"error": "Not authorized to update this post"})
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	models.DB.Save(&post)

	c.JSON(200, post)
}

func DeletePostById(c *gin.Context) {
	postId := c.Param("id")
	userIDValue, exists := c.Get("userID")

	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	userId := userIDValue.(uint)

	var post models.Post
	result := models.DB.First(&post, postId)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	if post.UserID != userId {
		c.JSON(403, gin.H{"error": "Not authorized to delete this post"})
		return
	}

	models.DB.Delete(&post)

	c.JSON(204, nil)
}

func CreatePost(c *gin.Context) {
	var post models.Post
	userIDValue, exists := c.Get("userID")

	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	userId := userIDValue.(uint)

	if err := c.BindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Set the post's user ID to the authenticated user
	post.UserID = userId

	models.DB.Create(&post)

	c.JSON(201, post)
}
