package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
	"github.com/risqiikhsani/rentvehicles/configs"
	"github.com/risqiikhsani/rentvehicles/middlewares"
	"github.com/risqiikhsani/rentvehicles/models"
	"github.com/risqiikhsani/rentvehicles/routes"
	socket "github.com/risqiikhsani/rentvehicles/socketio"
	"github.com/risqiikhsani/rentvehicles/utils"
	"github.com/risqiikhsani/rentvehicles/websocket"
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
	// f, _ := os.Create(appConfig.LogFile)
	//gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	middlewares.InitializeLogging(appConfig.LogFile)
	// logrus.SetOutput(io.MultiWriter(f, os.Stdout))

	// utils.InitializeValidator()
	r := gin.Default()

	// CORS middleware to allow requests from localhost:3000 (Next.js development)
	r.Use(func(c *gin.Context) {
		// // List of allowed origins (add your URLs here)
		// allowedOrigins := []string{
		// 	"http://localhost:3000",
		// 	"http://192.168.1.3:3000",
		// 	// Add more URLs as needed
		// }

		// origin := c.GetHeader("Origin")
		// for _, allowedOrigin := range allowedOrigins {
		// 	if allowedOrigin == origin {
		// 		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		// 		break
		// 	}
		// }

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins

		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Include DELETE here
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

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
	// public.Use(cors.Default())
	// public.Use(middlewares.AuthMiddleware())
	// public.Use(middlewares.LogMiddleware())
	routes.SetupPublicAccountRoutes(public)
	routes.SetupUserRoutes(public)
	routes.SetupMeRoutes(public)
	routes.SetupPostRoutes(public)
	routes.SetupFavoriteRoutes(public)
	routes.SetupLocationRoutes(public)
	routes.SetupCatRoutes(public, dbInstance)
	routes.SetupRentRoutes(public)
	routes.SetupRentDetailRoutes(public)
	// r.GET("/websocket", websocket.Ws)
	// r.GET("/websocket", func(c *gin.Context) {
	// 	websocket.ServeWs(c.Writer, c.Request)
	// })

	// go websocket.BroadcastMessages("Test")

	sockets := r.Group("/socket")

	var hub = websocket.NewHub() // create new instance of hub
	go hub.Run()                 // starting go routine

	// define end point where websocket is
	sockets.GET("/websocket", func(c *gin.Context) {
		websocket.ServeWs(hub, c.Writer, c.Request)
	})

	socket.InitializeSocketIO(sockets)

	addr := fmt.Sprintf(":%s", serverPort)

	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://google.com"}
	// // config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	// config.AllowAllOrigins = true

	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
