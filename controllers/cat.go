package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/handlers"
	"github.com/risqiikhsani/rentvehicles/models"
	"github.com/risqiikhsani/rentvehicles/utils"
)

func GetCats(db models.CatDatabase, auth handlers.Authenticator) gin.HandlerFunc {
	return func(c *gin.Context) {
		cats, err := db.GetCats()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cats not found"})
			return
		}
		c.JSON(http.StatusOK, cats)
	}
}

func GetCatById(db models.CatDatabase, auth handlers.Authenticator) gin.HandlerFunc {
	return func(c *gin.Context) {
		catID := c.Param("cat_id")
		cat, err := db.GetCatByID(catID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cat not found"})
			return
		}
		c.JSON(http.StatusOK, cat)
	}
}

func CreateCat(db models.CatDatabase, auth handlers.Authenticator) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _, authenticated := auth.RequireAuthentication2(c, "")
		if !authenticated {
			return
		}

		var cat models.Cat
		if err := c.ShouldBindJSON(&cat); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := utils.Validate.Struct(cat); err != nil {
			errs := utils.TranslateError(err, utils.En)
			c.JSON(400, gin.H{"errors": errs})
			return
		}

		cat.UserID = userID
		// Create the cat in the database
		err := db.CreateCat(&cat)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create cat"})
			return
		}

		c.JSON(201, cat)
	}
}

func UpdateCatById(db models.CatDatabase, auth handlers.Authenticator) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user is authenticated
		userID, _, authenticated := auth.RequireAuthentication2(c, "")
		if !authenticated {
			return
		}

		catId := c.Param("cat_id")

		existingCat, err := db.GetCatByID(catId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cat not found"})
			return
		}
		// Check if the user is the owner of the cat
		if existingCat.UserID != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this cat"})
			return
		}

		if err := c.ShouldBindJSON(&existingCat); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := utils.Validate.Struct(existingCat); err != nil {
			errs := utils.TranslateError(err, utils.En)
			c.JSON(400, gin.H{"errors": errs})
			return
		}

		err = db.UpdateCat(existingCat)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cat"})
			return
		}

		c.JSON(http.StatusOK, existingCat)

	}
}

func DeleteCatById(db models.CatDatabase, auth handlers.Authenticator) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user is authenticated
		userID, _, authenticated := auth.RequireAuthentication2(c, "")
		if !authenticated {
			return
		}

		catId := c.Param("cat_id")

		existingCat, err := db.GetCatByID(catId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cat not found"})
			return
		}

		if existingCat.UserID != userID {
			c.JSON(403, gin.H{"error": "Not authorized to delete this cat"})
			return
		}

		err = db.DeleteCat(existingCat)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete cat"})
			return
		}

		c.JSON(204, nil)
	}
}
