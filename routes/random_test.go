package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPing(t *testing.T) {
	// Create a Gin router and a recorder for the response
	// router := gin.Default()
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/ping", nil)

	// Create a context with the request and recorder
	c, _ := gin.CreateTestContext(recorder)
	c.Request = req

	// Call the handler function
	Ping(c)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, recorder.Code)
	}

	// Check the response body
	expectedBody := `{"message":"pong"}`
	if recorder.Body.String() != expectedBody {
		t.Errorf("Expected body %s; got %s", expectedBody, recorder.Body.String())
	}
}
