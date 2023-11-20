package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/handlers"
	"github.com/risqiikhsani/rentvehicles/models"
	"github.com/risqiikhsani/rentvehicles/utils"
)

// Implement other route handlers similarly
func GetUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(200, users)
}

func GetUserById(c *gin.Context) {
	userID := c.Param("user_id")

	var user models.User

	// Find the user by ID
	result := models.DB.First(&user, userID)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

func GetMeUser(c *gin.Context) {
	// Check if the user is authenticated
	userID, _, authenticated := handlers.RequireAuthentication(c, "")
	if !authenticated {
		return
	}

	var user models.User

	// Find the user by ID
	result := models.DB.First(&user, userID)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

func UpdateMeUser(c *gin.Context) {
	// Check if the user is authenticated
	userID, _, authenticated := handlers.RequireAuthentication(c, "")
	if !authenticated {
		return
	}

	// Check if the rent exists
	var existingUser models.User
	if err := models.DB.First(&existingUser, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// only allow rent maker to update user data
	if existingUser.ID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this user"})
		return
	}

	// Update the text of the user
	if err := c.ShouldBindJSON(&existingUser); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	if err := utils.Validate.Struct(existingUser); err != nil {
		errs := utils.TranslateError(err, utils.En)
		c.JSON(400, gin.H{"errors": errs})
		return
	}

	if err := models.DB.Save(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingUser)
}

// func UpdateUserById(c *gin.Context) {
// 	var user models.User
// 	userId := c.Param("user_id")

// 	result := models.DB.First(&user, userId)
// 	if result.Error != nil {
// 		c.JSON(404, gin.H{"error": "User not found"})
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	models.DB.Save(&user)

// 	c.JSON(200, user)
// }

// func DeleteUserById(c *gin.Context) {
// 	userId := c.Param("user_id")

// 	var user models.User
// 	result := models.DB.First(&user, userId)
// 	if result.Error != nil {
// 		c.JSON(404, gin.H{"error": "User not found"})
// 		return
// 	}

// 	models.DB.Delete(&user)
// 	c.JSON(204, nil)
// }

// func CreateUser(c *gin.Context) {
// 	var user models.User
// 	if err := c.BindJSON(&user); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}
// 	models.DB.Create(&user)
// 	c.JSON(201, user)
// }
