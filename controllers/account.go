package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/handlers"
	"github.com/risqiikhsani/rentvehicles/models"
	"github.com/risqiikhsani/rentvehicles/utils"
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

	var existingUser models.User
	if err := models.DB.Where("account_id = ?", acc.ID).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User data not found"})
		return
	}

	c.JSON(201, existingUser)
}

func RegisterAdmin(c *gin.Context) {
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
		Role:           "admin",
	}
	models.DB.Create(&usr)

	var existingUser models.User
	if err := models.DB.Where("account_id = ?", acc.ID).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User data not found"})
		return
	}

	c.JSON(201, existingUser)
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

	c.JSON(http.StatusOK, gin.H{"token": token, "user": existingUser})

}

type AccountUpdate struct {
	Email                   string `json:"email"`
	Password                string `json:"password"`
	Phone                   string `json:"phone"`
	ValidateCurrentPassword string `json:"validate_current_password" binding:"required"`
}

func UpdateAccount(c *gin.Context) {
	// Check if the user is authenticated
	userID, _, authenticated := handlers.CheckAuthentication(c)
	if !authenticated {
		return
	}

	var existingUser models.User
	if err := models.DB.First(&existingUser, userID).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// check if account is owner's
	var existingAccount models.Account
	if err := models.DB.First(&existingAccount, existingUser.AccountID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Account not found"})
		return
	}

	// Create a new instance of AccountUpdate and bind JSON data to it
	var accountUpdate AccountUpdate
	if err := c.ShouldBindJSON(&accountUpdate); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingAccount.Password), []byte(accountUpdate.ValidateCurrentPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Current Password is wrong"})
		return
	}

	// Copy the allowed fields from accountUpdate to existingAccount
	// Only update the fields if they are provided in the request
	if accountUpdate.Email != "" {
		existingAccount.Email = accountUpdate.Email
	}

	if accountUpdate.Password != "" {
		existingAccount.Password = accountUpdate.Password
	}

	if accountUpdate.Phone != "" {
		existingAccount.Phone = accountUpdate.Phone
	}

	models.DB.Save(&existingAccount)

	c.JSON(200, existingAccount)
}

func GetAccount(c *gin.Context) {
	// Check if the user is authenticated
	userID, _, authenticated := handlers.CheckAuthentication(c)
	if !authenticated {
		return
	}

	var existingUser models.User
	if err := models.DB.First(&existingUser, userID).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// check if account is owner's
	var existingAccount models.Account
	if err := models.DB.First(&existingAccount, existingUser.AccountID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Account not found"})
		return
	}

	c.JSON(200, existingAccount)
}
