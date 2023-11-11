package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/controllers"
	"github.com/risqiikhsani/rentvehicles/middlewares"
)

func SetupRentRoutes(public *gin.RouterGroup) {
	rentGroup := public.Group("/rents")
	rentGroup.Use(middlewares.AuthMiddleware())
	rentGroup.Use(middlewares.LogMiddleware())
	{
		rentGroup.GET("estimateprice", controllers.GetEstimateRentPrice)
		rentGroup.GET("", controllers.GetRents)
		rentGroup.GET("/:rent_id", controllers.GetRentById)
		rentGroup.PUT("/:rent_id", controllers.UpdateRentById)
		// rents data intends to be not deleted , remains history, only can be cancelled
		// rentGroup.DELETE("/:rent_id", controllers.DeletePostById)
		rentGroup.POST("", controllers.CreateRent)
		// if it's paid , can't cancel the rent ? or refundable rent
		rentGroup.POST("/:rent_id/cancel", nil)
	}
}
