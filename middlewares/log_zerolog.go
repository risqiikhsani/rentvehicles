package middlewares

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logFile *os.File

func InitializeLogging(logFilePath string) {
	logFile = createLogFile(logFilePath)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// output to console
	// log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// output to file
	// log.Logger = zerolog.New(logFile).With().Timestamp().Logger()

	// Create a multiwriter to write to both the file and os.Stderr (console)
	multi := zerolog.MultiLevelWriter(logFile, zerolog.ConsoleWriter{Out: os.Stderr})

	// Set up the logger to output to the multiwriter and include timestamps
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()
}

func createLogFile(logFilePath string) *os.File {
	f, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("error opening file")
	}
	return f
}

// LogMiddleware is responsible for logging requests and responses
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Log the request
		userID, _ := c.Get("userID")
		clientIP := c.ClientIP()
		requestHeaders := filterSensitiveHeaders(c.Request.Header)
		httpMethod := c.Request.Method
		httpPath := c.Request.URL.Path

		// Get the PID
		pid := os.Getpid()

		log.Info().
			Int("pid", pid). // Include the PID as a field in the log
			Str("method", httpMethod).
			Str("path", httpPath).
			Interface("userID", userID).
			Str("clientIP", clientIP).
			Interface("requestHeaders", requestHeaders).
			Msg("Request received")

		// Continue handling the request
		c.Next()

		// Log the response and request duration
		latency := time.Since(start)

		responseHeaders := c.Writer.Header()
		httpStatus := c.Writer.Status()
		responsePayload := c.Writer.Size()

		log.Info().
			Int("pid", pid). // Include the PID in the response log as well
			Str("method", httpMethod).
			Str("path", httpPath).
			Int("status", httpStatus).
			Dur("duration", latency).
			Interface("userID", userID).
			Str("clientIP", clientIP).
			Interface("responseHeaders", responseHeaders).
			Int("responsePayload", responsePayload).
			Msg("Request completed")
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
