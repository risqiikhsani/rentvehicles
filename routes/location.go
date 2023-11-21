package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/controllers"
	"github.com/risqiikhsani/rentvehicles/middlewares"
)

func SetupLocationRoutes(public *gin.RouterGroup) {
	locationGroup := public.Group("/locations")
	{
		locationGroup.GET("", controllers.GetLocations)                 // done test
		locationGroup.GET("/:location_id", controllers.GetLocationById) // done test
	}
	locationGroup.Use(middlewares.AuthMiddleware())
	locationGroup.Use(middlewares.LogMiddleware())
	{
		locationGroup.POST("", controllers.CreateLocation)                    // done test
		locationGroup.PUT("/:location_id", controllers.UpdateLocationById)    // done test
		locationGroup.DELETE("/:location_id", controllers.DeleteLocationById) //done test
	}
}
