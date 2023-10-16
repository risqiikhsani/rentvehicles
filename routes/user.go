package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/controllers"
	"github.com/risqiikhsani/rentvehicles/middlewares"
)

func SetupUserRoutes(public *gin.RouterGroup) {
	userGroup := public.Group("/users")
	userGroup.Use(middlewares.AuthMiddleware())
	{
		userGroup.GET("", controllers.GetUsers)
		userGroup.GET("/:user_id", controllers.GetUserById)
		userGroup.PUT("/:user_id", controllers.UpdateUserById)
		userGroup.DELETE("/:user_id", controllers.DeleteUserById)
		userGroup.POST("", controllers.CreateUser)
	}
}
