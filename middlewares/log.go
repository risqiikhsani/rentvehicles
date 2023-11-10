package middlewares

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logFile *os.File

func InitializeLogging(logFilePath string) {
	logFile = createLogFile(logFilePath)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(logFile)
}

func createLogFile(logFilePath string) *os.File {
	f, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}
	return f
}

// LogMiddleware is responsible for logging requests and responses
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Log the request
		userID, _ := c.Get("userID")
		clientIP := c.Request.RemoteAddr
		requestHeaders := filterSensitiveHeaders(c.Request.Header)
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

// Function to filter sensitive headers
func filterSensitiveHeaders(headers map[string][]string) map[string][]string {
	filteredHeaders := make(map[string][]string)

	for key, values := range headers {
		if key == "Authorization" {
			filteredHeaders[key] = []string{"[REDACTED]"} // Redact sensitive header
		} else {
			filteredHeaders[key] = values
		}
	}

	return filteredHeaders
}
