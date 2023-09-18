package handlers

import (
	"github.com/gin-gonic/gin"
)

func CheckAuthentication(c *gin.Context) (uint, bool) {
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(404, gin.H{"error": "Unauthorized"})
		return 0, false
	}
	userID := userIDValue.(uint)
	return userID, true
}
