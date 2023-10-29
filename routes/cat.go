package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/controllers"

	"github.com/risqiikhsani/rentvehicles/middlewares"
	"github.com/risqiikhsani/rentvehicles/models"
)

var db models.CatDatabase

func SetupCatRoutes(public *gin.RouterGroup, dbinstance *models.MyDatabase) {

	db = dbinstance

	catGroup := public.Group("/cats")
	catGroup.Use(middlewares.AuthMiddleware())
	{
		catGroup.POST("", controllers.CreateCat(db))               // done test
		catGroup.GET("", controllers.GetCats(db))                  // done test
		catGroup.GET("/:cat_id", controllers.GetCatById(db))       // done test
		catGroup.PUT("/:cat_id", controllers.UpdateCatById(db))    // done test
		catGroup.DELETE("/:cat_id", controllers.DeleteCatById(db)) //done test
	}
}
