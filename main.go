package main

import (
	"fmt"
	"io"
	"log"
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
	if err != nil {
		panic(err)
	}

	secretConfig, err := configs.LoadSecretConfig("./")
	if err != nil {
		panic(err)
	}

	configs.SetMainConfig(appConfig)
	configs.SetSecretConfig(secretConfig)

	fmt.Println("JWT " + secretConfig.SecretKey)

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

	// Create an instance of the database
	dbInstance, err := models.ConnectDB(secretConfig) // Use your models package function
	if err != nil {
		// Handle the error
		panic(err)
	}

	public := r.Group("/api")
	public.Use(middlewares.AuthMiddleware())
	public.Use(middlewares.LogMiddleware())
	routes.SetupPublicAccountRoutes(public)
	routes.SetupUserRoutes(public)
	routes.SetupAccountRoutes(public)
	routes.SetupPostRoutes(public)
	routes.SetupLocationRoutes(public)
	routes.SetupCatRoutes(public, dbInstance)
	routes.SetupRentRoutes(public)

	addr := fmt.Sprintf(":%s", serverPort)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
