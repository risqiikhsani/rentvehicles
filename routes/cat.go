package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/controllers"
	"github.com/risqiikhsani/rentvehicles/handlers"

	"github.com/risqiikhsani/rentvehicles/middlewares"
	"github.com/risqiikhsani/rentvehicles/models"
)

var db models.CatDatabase
var auth handlers.Authenticator

func SetupCatRoutes(public *gin.RouterGroup, dbinstance *models.MyDatabase) {

	db = dbinstance
	auth = &handlers.DefaultAuthenticator{}

	catGroup := public.Group("/cats")
	catGroup.Use(middlewares.AuthMiddleware())
	catGroup.Use(middlewares.LogMiddleware())
	{
		catGroup.POST("", controllers.CreateCat(db, auth))               // done test
		catGroup.GET("", controllers.GetCats(db, auth))                  // done test
		catGroup.GET("/:cat_id", controllers.GetCatById(db, auth))       // done test
		catGroup.PUT("/:cat_id", controllers.UpdateCatById(db, auth))    // done test
		catGroup.DELETE("/:cat_id", controllers.DeleteCatById(db, auth)) //done test
	}
}
