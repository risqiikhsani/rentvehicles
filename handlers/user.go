package handlers

import (
	"github.com/gin-gonic/gin"
)

func CheckAuthentication(c *gin.Context) (uint, string, bool) {

	userIDValue, a := c.Get("userID")
	userRoleValue, b := c.Get("userRole") // Assuming you store the user's role in "userRole"
	authenticated := false

	if !a || !b {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return 0, "", false
	}

	userID := userIDValue.(uint)
	userRole := userRoleValue.(string)
	authenticated = true

	return userID, userRole, authenticated
}
