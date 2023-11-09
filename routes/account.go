package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/controllers"
	"github.com/risqiikhsani/rentvehicles/middlewares"
)

func SetupPublicAccountRoutes(public *gin.RouterGroup) {
	public.POST("/register", controllers.Register)            // done test
	public.POST("/register-admin", controllers.RegisterAdmin) // done test
	public.POST("/login", controllers.Login)                  // done test

}

func SetupAccountRoutes(public *gin.RouterGroup) {
	accountGroup := public.Group("/account")
	accountGroup.Use(middlewares.AuthMiddleware())
	accountGroup.Use(middlewares.LogMiddleware())
	{
		public.PUT("", controllers.UpdateAccount) // done test
		public.GET("", controllers.GetAccount)    // done test
	}
}
