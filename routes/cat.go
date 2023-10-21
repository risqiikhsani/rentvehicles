package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/controllers"
	"github.com/risqiikhsani/rentvehicles/middlewares"
)

func SetupCatRoutes(public *gin.RouterGroup) {
	catGroup := public.Group("/cats")
	catGroup.Use(middlewares.AuthMiddleware())
	{
		catGroup.POST("", controllers.CreateCat)               // done test
		catGroup.GET("", controllers.GetCats)                  // done test
		catGroup.GET("/:cat_id", controllers.GetCatById)       // done test
		catGroup.PUT("/:cat_id", controllers.UpdateCatById)    // done test
		catGroup.DELETE("/:cat_id", controllers.DeleteCatById) //done test
	}
}
