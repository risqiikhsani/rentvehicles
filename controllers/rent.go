package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/handlers"
	"github.com/risqiikhsani/rentvehicles/models"
	"github.com/risqiikhsani/rentvehicles/utils"
	"gorm.io/gorm/clause"
)

func GetEstimateRentPrice(c *gin.Context) {
	// Check if the user is authenticated
	_, _, authenticated := handlers.RequireAuthentication(c, "")
	if !authenticated {
		return
	}

	// Get the parameters from the request
	postIDstr := c.Query("post_id")
	rentDaysStr := c.DefaultQuery("rent_days", "1")
	voucherCodeStr := c.DefaultQuery("voucher_code", "") // Voucher code is optional
	postId, errPostId := strconv.ParseUint(postIDstr, 10, 32)
	rentDays, errRentDays := strconv.ParseUint(rentDaysStr, 10, 32)

	if errPostId != nil {
		// Handle the error for postId conversion
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
		return
	}

	if errRentDays != nil {
		// Handle the error for rentDays conversion
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rent days"})
		return
	}

	// Calculate the estimated rental price based on the parameters
	// Replace the following line with your calculation logic

	var post models.Post

	if err := models.DB.Preload(clause.Associations).First(&post, postId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	finalPrice, normalPrice, savedPrice, _ := post.CalculateRentalPrice(uint(rentDays))

	c.JSON(http.StatusOK, gin.H{
		"post_id":                postId,
		"rent_days":              rentDays,
		"voucher_code":           voucherCodeStr,
		"estimated_final_price":  finalPrice,
		"estimated_normal_price": normalPrice,
		"estimated_saved_price":  savedPrice,
	})
}

func GetRents(c *gin.Context) {
	userID, userRole, authenticated := handlers.RequireAuthentication(c, "")
	if !authenticated {
		return
	}

	var rents []models.Rent

	// if basic user , rents data will be the user's rents history
	if userRole == "Basic" {
		models.DB.Preload(clause.Associations).Where("user_id = ?", userID).Find(&rents)
		// if admin user, rents data will be the rents data which post_id is admin's
	} else if userRole == "Admin" {
		// Assuming an admin can only see rents associated with their own posts.
		subquery := models.DB.Model(&models.Post{}).Select("ID").Where("user_id = ?", userID)
		models.DB.Preload(clause.Associations).Where("post_id IN (?)", subquery).Find(&rents)
	} else {
		c.JSON(http.StatusForbidden, gin.H{"message": "Permission denied"})
		return
	}

	c.JSON(http.StatusOK, rents)
}

func GetRentById(c *gin.Context) {
	_, _, authenticated := handlers.RequireAuthentication(c, "")
	if !authenticated {
		return
	}

	rent_id := c.Param("rent_id")

	var rent models.Rent

	result := models.DB.Preload(clause.Associations).First(&rent, rent_id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Rent not found"})
		return
	}

	c.JSON(200, rent)
}

func CreateRent(c *gin.Context) {
	userID, _, authenticated := handlers.RequireAuthentication(c, "Basic")
	if !authenticated {
		return
	}

	var rent models.Rent
	if err := c.ShouldBindJSON(&rent); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	rent.UserID = userID

	if err := utils.Validate.Struct(rent); err != nil {
		errs := utils.TranslateError(err, utils.En)
		c.JSON(400, gin.H{"errors": errs})
		return
	}

	// do some checking in rent's hook

	// Create the rent in the database
	if err := models.DB.Create(&rent).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, rent)

}

func UpdateRentById(c *gin.Context) {
	// Check if the user is authenticated
	userID, _, authenticated := handlers.RequireAuthentication(c, "Basic")
	if !authenticated {
		return
	}

	// Get the rent ID from the URL parameters
	rent_id := c.Param("rent_id")

	// Check if the rent exists
	var existingRent models.Rent
	if err := models.DB.Where("id = ?", rent_id).First(&existingRent).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rent not found"})
		return
	}

	// return error if associated post is not found
	var associatedPost models.Post
	if err := models.DB.First(&associatedPost, existingRent.PostID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Associated post not found"})
		return
	}

	// only allow rent maker to update rent data
	if existingRent.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this rent"})
		return
	}

	// Update the text of the rent
	if err := c.ShouldBindJSON(&existingRent); err != nil {
		c.JSON(400, gin.H{"errors": err.Error()})
		return
	}

	if err := utils.Validate.Struct(existingRent); err != nil {
		errs := utils.TranslateError(err, utils.En)
		c.JSON(400, gin.H{"errors": errs})
		return
	}

	if err := models.DB.Save(&existingRent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingRent)
}
