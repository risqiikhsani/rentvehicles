package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/controllers"
	"github.com/risqiikhsani/rentvehicles/middlewares"
)

func SetupPostRoutes(public *gin.RouterGroup) {
	postGroup := public.Group("/posts")
	{
		postGroup.GET("", controllers.GetPosts)             // done test
		postGroup.GET("/:post_id", controllers.GetPostById) // done test
	}
	postGroup.Use(middlewares.AuthMiddleware())
	postGroup.Use(middlewares.LogMiddleware())
	{
		postGroup.GET("/created", controllers.GetMyPosts)
		postGroup.PUT("/:post_id", controllers.UpdatePostById)    // done test
		postGroup.DELETE("/:post_id", controllers.DeletePostById) // done test
		postGroup.POST("", controllers.CreatePost)                // done test
	}
}
