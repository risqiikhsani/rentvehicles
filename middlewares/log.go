package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LogMiddleware is responsible for logging requests and responses
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Log the request
		logrus.WithFields(logrus.Fields{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		}).Info("Request received")

		// Continue handling the request
		c.Next()

		// Log the response and request duration
		latency := time.Since(start)
		logrus.WithFields(logrus.Fields{
			"method":   c.Request.Method,
			"path":     c.Request.URL.Path,
			"status":   c.Writer.Status(),
			"duration": latency,
		}).Info("Request completed")
	}
}
