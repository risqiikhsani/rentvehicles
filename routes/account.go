package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/controllers"
)

func SetupPublicAccountRoutes(public *gin.RouterGroup) {
	public.POST("/register", controllers.Register)
	public.POST("/register-admin", controllers.RegisterAdmin)
	public.POST("/login", controllers.Login)

}

func SetupAccountRoutes(public *gin.RouterGroup) {
	public.PUT("/account", controllers.UpdateAccount)
	public.GET("/account", controllers.GetAccount)
}
