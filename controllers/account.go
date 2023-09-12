package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/contactgo/models"
	"github.com/risqiikhsani/contactgo/utils"
	"golang.org/x/crypto/bcrypt"
)

// type UpdateInput struct {
// 	Email     string `json:"email"`
// 	Password  string `json:"password"`
// }

// func Update(c *gin.Context){
// 	var input UpdateInput

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(400,gin.H{"error":err.Error()})
// 		return
// 	}

// }

type RegisterInput struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	Password2 string `json:"password2" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Check if the passwords match
	if input.Password != input.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	// Check if the account already exists with the same username
	var existingAccount models.Account
	if err := models.DB.Where("username = ?", input.Username).First(&existingAccount).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// Hash the password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash the password"})
		return
	}

	acc := models.Account{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}
	// create the account
	models.DB.Create(&acc)

	// if success, automatically create user associated, or user manually create it
	// Generate a random string for PublicUsername
	publicUsername := utils.RandomString()
	usr := models.User{
		PublicUsername: publicUsername,
		AccountID:      acc.ID, // Assuming AccountID is of uint type
	}
	models.DB.Create(&usr)

	c.JSON(201, acc)
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// check if username not exists
	var existingAccount models.Account
	if err := models.DB.Where("username = ?", input.Username).First(&existingAccount).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username not exists"})
		return
	}

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(existingAccount.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// get user data
	var existingUser models.User
	if err := models.DB.Where("account_id = ?", existingAccount.ID).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User data not found"})
		return
	}

	// Generate and return a JWT token on successful login
	token, err := utils.GenerateJWTToken(existingUser.ID, existingUser.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
