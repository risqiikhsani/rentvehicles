package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/controllers"
	"github.com/risqiikhsani/rentvehicles/middlewares"
)

func SetupPaymentRoutes(public *gin.RouterGroup) {
	postGroup := public.Group("/payments")
	postGroup.Use(middlewares.AuthMiddleware())
	postGroup.Use(middlewares.LogMiddleware())
	{
		postGroup.GET("/create", controllers.CreatePayment)
	}
}
