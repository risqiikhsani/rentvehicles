package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/contactgo/controllers"
	"github.com/risqiikhsani/contactgo/middlewares"
)

func SetupPostRoutes(public *gin.RouterGroup) {
	postGroup := public.Group("/posts")
	postGroup.Use(middlewares.AuthMiddleware())
	{
		postGroup.GET("", controllers.GetPosts)
		postGroup.GET("/:id", controllers.GetPostById)
		postGroup.PUT("/:id", controllers.UpdatePostById)
		postGroup.DELETE("/:id", controllers.DeletePostById)
		postGroup.POST("", controllers.CreatePost)
	}
}
