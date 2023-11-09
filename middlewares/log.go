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
		userID, _ := c.Get("userID")
		// clientIP := c.ClientIP() // Get client's IP address
		clientIP := c.Request.RemoteAddr
		requestHeaders := c.Request.Header
		httpMethod := c.Request.Method
		httpPath := c.Request.URL.Path

		logrus.WithFields(logrus.Fields{
			"method":         httpMethod,
			"path":           httpPath,
			"userID":         userID,
			"clientIP":       clientIP,
			"requestHeaders": requestHeaders,
		}).Info("Request received")

		// Continue handling the request
		c.Next()

		// Log the response and request duration
		latency := time.Since(start)

		responseHeaders := c.Writer.Header()
		httpStatus := c.Writer.Status()
		responsePayload := c.Writer.Size()

		logrus.WithFields(logrus.Fields{
			"method":          httpMethod,
			"path":            httpPath,
			"status":          httpStatus,
			"duration":        latency,
			"userID":          userID,
			"clientIP":        clientIP,
			"responseHeaders": responseHeaders,
			"responsePayload": responsePayload,
		}).Info("Request completed")
	}
}
