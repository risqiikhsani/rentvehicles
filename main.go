package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/contactgo/middlewares"
	"github.com/risqiikhsani/contactgo/models"
	"github.com/risqiikhsani/contactgo/routes"
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

	// Force log's color
	// gin.ForceConsoleColor()

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create(viper.GetString("log_file"))
	//gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	models.ConnectDB()

	public := r.Group("/api")
	public.Use(middlewares.LogMiddleware())
	routes.SetupUserRoutes(public)
	routes.SetupAccountRoutes(public)
	routes.SetupPostRoutes(public)

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
