package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/controllers"
	"github.com/risqiikhsani/rentvehicles/middlewares"
)

func SetupLocationRoutes(public *gin.RouterGroup) {
	locationGroup := public.Group("/locations")
	locationGroup.Use(middlewares.AuthMiddleware())
	{
		locationGroup.GET("", controllers.GetLocations)
		locationGroup.GET("/:location_id", controllers.GetLocationById)
		locationGroup.PUT("/:location_id", controllers.UpdateLocationById)
		locationGroup.DELETE("/:location_id", controllers.DeleteLocationById)
		locationGroup.POST("", controllers.CreateLocation)
	}
}
