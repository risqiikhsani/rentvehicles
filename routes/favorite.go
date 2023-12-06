package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/controllers"
	"github.com/risqiikhsani/rentvehicles/middlewares"
)

func SetupFavoriteRoutes(public *gin.RouterGroup) {
	me := public.Group("/favorite")
	me.Use(middlewares.AuthMiddleware())
	me.Use(middlewares.LogMiddleware())
	{
		me.GET("/", controllers.GetFavorites)
		me.POST("/", controllers.CreateFavorite)
		me.DELETE("/", controllers.DeleteFavoriteById)
	}
}
