package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/contactgo/controllers"
)

func SetupUserRoutes(public *gin.RouterGroup) {
	userGroup := public.Group("/users")
	{
		userGroup.GET("", controllers.GetUsers)
		userGroup.GET("/:id", controllers.GetUserById)
		userGroup.PUT("/:id", controllers.UpdateUserById)
		userGroup.DELETE("/:id", controllers.DeleteUserById)
		userGroup.POST("", controllers.CreateUser)
	}
}
