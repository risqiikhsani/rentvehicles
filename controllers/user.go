package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/contactgo/models"
)

// Implement other route handlers similarly
func GetUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(200, users)
}

func GetUserById(c *gin.Context) {
	userID := c.Param("id")

	var user models.User

	// Find the user by ID
	result := models.DB.First(&user, userID)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

func UpdateUserById(c *gin.Context) {
	var user models.User
	userId := c.Param("id")

	result := models.DB.First(&user, userId)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	models.DB.Save(&user)

	c.JSON(200, user)
}

func DeleteUserById(c *gin.Context) {
	userId := c.Param("id")

	var user models.User
	result := models.DB.First(&user, userId)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	models.DB.Delete(&user)
	c.JSON(204, nil)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	models.DB.Create(&user)
	c.JSON(201, user)
}
