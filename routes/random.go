package routes

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var randomGroup *gin.Engine

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func SetupTestRoutes(public *gin.RouterGroup) {
	randomGroup := public.Group("/random-test")
	{
		randomGroup.GET("/cookie", func(c *gin.Context) {

			cookie, err := c.Cookie("gin_cookie")

			if err != nil {
				cookie = "NotSet"
				c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
			}

			fmt.Printf("Cookie value: %s \n", cookie)
		})

		randomGroup.GET("/ping", Ping)

		randomGroup.GET("/long_async", func(c *gin.Context) {
			// create copy to be used inside the goroutine
			cCp := c.Copy()
			go func() {
				// simulate a long task with time.Sleep(). 5 seconds
				time.Sleep(5 * time.Second)

				// note that you are using the copied context "cCp", IMPORTANT
				log.Println("Done! in path " + cCp.Request.URL.Path)
			}()
		})

		randomGroup.GET("/long_sync", func(c *gin.Context) {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// since we are NOT using a goroutine, we do not have to copy the context
			log.Println("Done! in path " + c.Request.URL.Path)
		})

		randomGroup.POST("/post", func(c *gin.Context) {
			id := c.Query("id")

			title := c.PostForm("title")
			text := c.PostForm("text")

			fmt.Printf("id: %s;title:%s;text:%s", id, title, text)
		})
	}
}
