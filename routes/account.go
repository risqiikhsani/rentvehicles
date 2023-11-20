package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/controllers"
)

func SetupPublicAccountRoutes(public *gin.RouterGroup) {
	public.POST("/register", controllers.Register)            // done test
	public.POST("/register-admin", controllers.RegisterAdmin) // done test
	public.POST("/login", controllers.Login)                  // done test
	public.POST("/forgot-password", controllers.ForgotPassword)
	public.POST("/reset-password", controllers.ResetPassword)
}
