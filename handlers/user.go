package handlers

import (
	"github.com/gin-gonic/gin"
)

type Authenticator interface {
	RequireAuthentication2(c *gin.Context, requiredRole string) (uint, string, bool)
	CheckAuthentication2(c *gin.Context) (uint, string, bool)
}

type DefaultAuthenticator struct{}

func (a *DefaultAuthenticator) CheckAuthentication2(c *gin.Context) (uint, string, bool) {
	userIDValue, ok1 := c.Get("userID")
	userRoleValue, ok2 := c.Get("userRole") // Assuming you store the user's role in "userRole"
	authenticated := false

	if !ok1 || !ok2 {
		return 0, "", false
	}

	userID := userIDValue.(uint)
	userRole := userRoleValue.(string)
	authenticated = true

	return userID, userRole, authenticated
}

func (a *DefaultAuthenticator) RequireAuthentication2(c *gin.Context, requiredRole string) (uint, string, bool) {
	userID, userRole, authenticated := a.CheckAuthentication2(c)
	if !authenticated {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return 0, "", false
	}

	if requiredRole != "" && userRole != requiredRole {
		c.JSON(401, gin.H{"error": "You do not have permission"})
		return 0, "", false
	}

	return userID, userRole, true
}

func CheckAuthentication(c *gin.Context) (uint, string, bool) {

	userIDValue, a := c.Get("userID")
	userRoleValue, b := c.Get("userRole") // Assuming you store the user's role in "userRole"
	authenticated := false

	if !a || !b {
		return 0, "", false
	}

	userID := userIDValue.(uint)
	userRole := userRoleValue.(string)
	authenticated = true

	return userID, userRole, authenticated
}

func RequireAuthentication(c *gin.Context, requiredRole string) (uint, string, bool) {
	userID, userRole, authenticated := CheckAuthentication(c)
	if !authenticated {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return 0, "", false
	}

	if requiredRole != "" && userRole != requiredRole {
		c.JSON(401, gin.H{"error": "You do not have permission"})
		return 0, "", false
	}

	return userID, userRole, true
}
