package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/controllers"
	"github.com/risqiikhsani/rentvehicles/middlewares"
)

func SetupRentRoutes(public *gin.RouterGroup) {
	rentGroup := public.Group("/rents")
	rentGroup.Use(middlewares.AuthMiddleware())
	{
		rentGroup.GET("", controllers.GetPosts)
		rentGroup.GET("/:rent_id", controllers.GetPostById)
		rentGroup.PUT("/:rent_id", controllers.UpdatePostById)
		// rents data intends to be not deleted , remains history, only can be cancelled
		// rentGroup.DELETE("/:rent_id", controllers.DeletePostById)
		rentGroup.POST("", controllers.CreatePost)
		// if it's paid , can't cancel the rent ? or refundable rent
		rentGroup.POST("/:rent_id/cancel", nil)
	}
}
