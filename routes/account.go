package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/controllers"
)

func SetupPublicAccountRoutes(public *gin.RouterGroup) {
	public.POST("/register", controllers.Register)            // done test
	public.POST("/register-admin", controllers.RegisterAdmin) // done test
	public.POST("/login", controllers.Login)                  // done test

}

func SetupAccountRoutes(public *gin.RouterGroup) {
	public.PUT("/account", controllers.UpdateAccount) // done test
	public.GET("/account", controllers.GetAccount)    // done test
}
