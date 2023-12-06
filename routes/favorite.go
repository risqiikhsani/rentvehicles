package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/controllers"
	"github.com/risqiikhsani/rentvehicles/middlewares"
)

func SetupFavoriteRoutes(public *gin.RouterGroup) {
	fav := public.Group("/favorites")
	fav.Use(middlewares.AuthMiddleware())
	fav.Use(middlewares.LogMiddleware())
	{
		fav.GET("", controllers.GetFavorites)
		fav.GET("/posts", controllers.GetFavoritePosts)
		fav.POST("", controllers.CreateFavorite)
		fav.DELETE("", controllers.DeleteFavoriteById)
	}
}
