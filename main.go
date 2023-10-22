package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
	"github.com/risqiikhsani/rentvehicles/configs"
	"github.com/risqiikhsani/rentvehicles/middlewares"
	"github.com/risqiikhsani/rentvehicles/models"
	"github.com/risqiikhsani/rentvehicles/routes"
	"github.com/risqiikhsani/rentvehicles/utils"
	// "github.com/spf13/viper"
)

func main() {

	appConfig, err := configs.LoadAppConfig("./configs")
	secretConfig, err := configs.LoadSecretConfig("./")
	if err != nil {
		panic(err)
	}

	serverPort := appConfig.ServerPort

	baseURL := secretConfig.BaseUrl
	static_image_path := appConfig.StaticImagesPath
	models.SetBaseURL(baseURL)
	models.SetStaticImagePath(static_image_path)

	// Force log's color
	// gin.ForceConsoleColor()

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create(appConfig.LogFile)
	//gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// utils.InitializeValidator()
	r := gin.Default()

	utils.InitializeTranslator() // translator first , because initializeValidator needs it
	utils.InitializeValidator()

	static_path := appConfig.StaticPath
	r.Static("/static", "./"+static_path)

	models.ConnectDB(secretConfig)

	public := r.Group("/api")
	public.Use(middlewares.LogMiddleware())
	routes.SetupPublicAccountRoutes(public)  // in front of auth middleware so it's not using auth middleware (jwt token not required)
	public.Use(middlewares.AuthMiddleware()) // will apply to all routes below
	routes.SetupUserRoutes(public)
	routes.SetupAccountRoutes(public)
	routes.SetupPostRoutes(public)
	routes.SetupLocationRoutes(public)
	routes.SetupCatRoutes(public)

	addr := fmt.Sprintf(":%s", serverPort)
	r.Run(addr)
}
