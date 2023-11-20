package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/controllers"
	"github.com/risqiikhsani/rentvehicles/middlewares"
)

func SetupMeRoutes(public *gin.RouterGroup) {
	me := public.Group("/me")
	me.Use(middlewares.AuthMiddleware())
	me.Use(middlewares.LogMiddleware())
	{
		me.PUT("/account", controllers.UpdateAccount) // done test
		me.GET("/account", controllers.GetAccount)    // done test
		me.PUT("/user", controllers.UpdateMeUser)
		me.GET("/user", controllers.GetMeUser)
	}
}
