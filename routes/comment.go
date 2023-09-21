package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/contactgo/controllers"
	"github.com/risqiikhsani/contactgo/middlewares"
)

func SetupCommentRoutes(public *gin.RouterGroup) {
	commentGroup := public.Group("/")
	commentGroup.Use(middlewares.AuthMiddleware())
	{
		commentGroup.GET("/posts/:post_id/comments", controllers.GetComments)
		commentGroup.GET("/comments/:comment_id", controllers.GetCommentById)
		commentGroup.PUT("/comments/:comment_id", controllers.UpdateCommentById)
		commentGroup.DELETE("/comments/:comment_id", controllers.DeleteCommentById)
		commentGroup.POST("/posts/:post_id/comments", controllers.CreateComment)
	}
}
