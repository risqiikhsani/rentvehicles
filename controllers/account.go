package controllers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/configs"
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

	fmt.Println("user id")
	fmt.Println(existingUser.ID)
	fmt.Println("user role")
	fmt.Println(existingUser.Role)
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
		c.JSON(401, gin.H{"error": "Unauthorized"})
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

	// Hash the password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(accountUpdate.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash the password"})
		return
	}

	// Copy the allowed fields from accountUpdate to existingAccount
	// Only update the fields if they are provided in the request
	if accountUpdate.Email != "" {
		existingAccount.Email = accountUpdate.Email
	}

	if accountUpdate.Password != "" {
		existingAccount.Password = string(hashedPassword)
	}

	if accountUpdate.Phone != "" {
		existingAccount.Phone = accountUpdate.Phone
	}

	models.DB.Save(&existingAccount)

	c.JSON(200, existingAccount)
}

func GetAccount(c *gin.Context) {
	// Check if the user is authenticated
	userID, _, authenticated := handlers.RequireAuthentication(c, "")
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

type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required,email"`
}

func ForgotPassword(c *gin.Context) {
	var input ForgotPasswordInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the email exists in the database
	var existingAccount models.Account
	if err := models.DB.Where("email = ?", input.Email).First(&existingAccount).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not associated with any account"})
		return
	}

	// Generate a random token for password reset
	resetToken := utils.RandomStringUuid() // You can implement a function to generate secure random strings

	// Store the reset token and its expiration timestamp in the database
	resetTokenRecord := models.ForgotPassword{
		Token:     resetToken,
		AccountID: existingAccount.ID,
	}
	models.DB.Create(&resetTokenRecord)

	// URL-encode the token
	encodedToken := url.QueryEscape(resetToken)

	// Construct the reset password link by appending the encoded token to the URL
	resetPasswordURL := "http://localhost:3000/auth/reset-password/?token=" + encodedToken

	// Send an email to the user with the resetPasswordURL
	secretConf := configs.GetSecretConfig()
	sender := utils.NewGmailSender(secretConf.EmailSenderName, secretConf.EmailSenderAddress, secretConf.EmailSenderPassword)

	subject := "Forgot Password"
	content := `
		<h1>Hello</h1>
		<p>Forgot password has been requested. Click the following link to reset your password:</p>
		<a href="` + resetPasswordURL + `">Reset Password</a>
	`
	to := []string{existingAccount.Email}

	err := sender.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong. Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset instructions have been sent to your email"})
}

type ResetPasswordInput struct {
	Password  string `json:"password" binding:"required"`
	Password2 string `json:"password2" binding:"required"`
}

func ResetPassword(c *gin.Context) {
	token := c.DefaultQuery("token", "")

	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing token"})
		return
	}

	// Check if the token exists in the database and is not expired
	var resetTokenRecord models.ForgotPassword
	if err := models.DB.Where("token = ? AND expired = false", token).First(&resetTokenRecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired token"})
		return
	}

	var input ResetPasswordInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the passwords match
	if input.Password != input.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	// Hash the password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash the password"})
		return
	}

	var existingAccount models.Account
	if err := models.DB.First(&existingAccount, resetTokenRecord.AccountID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account not found"})
		return
	}

	existingAccount.Password = string(hashedPassword)

	// Update the account's password in the database
	if err := models.DB.Save(&existingAccount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password has been successfully updated."})
}
