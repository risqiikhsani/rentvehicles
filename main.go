package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/risqiikhsani/rentvehicles/middlewares"
	"github.com/risqiikhsani/rentvehicles/models"
	"github.com/risqiikhsani/rentvehicles/routes"
	"github.com/risqiikhsani/rentvehicles/utils"
	"github.com/spf13/viper"
)

func main() {

	viper.AddConfigPath("./configs")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("yaml")   // Look for specific type
	viper.ReadInConfig()
	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	serverPort := viper.GetString("server_port")

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	baseURL := os.Getenv("BASE_URL")
	static_image_path := viper.GetString("static_images_path")
	models.SetBaseURL(baseURL)
	models.SetStaticImagePath(static_image_path)

	// Force log's color
	// gin.ForceConsoleColor()

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create(viper.GetString("log_file"))
	//gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// utils.InitializeValidator()
	r := gin.Default()

	utils.InitializeTranslator() // translator first , because initializeValidator needs it
	utils.InitializeValidator()

	static_path := viper.GetString("static_path")
	r.Static("/static", "./"+static_path)

	models.ConnectDB()

	public := r.Group("/api")
	public.Use(middlewares.LogMiddleware())
	routes.SetupPublicAccountRoutes(public)  // in front of auth middleware so it's not using auth middleware (jwt token not required)
	public.Use(middlewares.AuthMiddleware()) // will apply to all routes below
	routes.SetupUserRoutes(public)
	routes.SetupAccountRoutes(public)
	routes.SetupPostRoutes(public)
	routes.SetupLocationRoutes(public)
	routes.SetupCatRoutes(public)

	public.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value: %s \n", cookie)
	})

	public.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	public.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// note that you are using the copied context "cCp", IMPORTANT
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	public.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(5 * time.Second)

		// since we are NOT using a goroutine, we do not have to copy the context
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	public.POST("/post", func(c *gin.Context) {
		id := c.Query("id")

		title := c.PostForm("title")
		text := c.PostForm("text")

		fmt.Printf("id: %s;title:%s;text:%s", id, title, text)
	})

	addr := fmt.Sprintf(":%s", serverPort)
	r.Run(addr)
}
