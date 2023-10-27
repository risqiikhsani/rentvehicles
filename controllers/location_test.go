package controllers

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	// import your models package
// )

// func TestGetLocations(t *testing.T) {
// 	// Initialize a Gin router and set its mode to TestMode
// 	r := gin.New()
// 	r.GET("/locations", GetLocations)

// 	// Create a mock database and replace the actual DB with it
// 	// You may want to use a package like github.com/DATA-DOG/go-sqlmock
// 	// to create a mock database for testing.
// 	// Replace models.DB with your mock database instance.

// 	// Perform a GET request to the /locations endpoint
// 	req, err := http.NewRequest("GET", "/locations", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	resp := httptest.NewRecorder()
// 	r.ServeHTTP(resp, req)

// 	// Check the response status code and body
// 	assert.Equal(t, http.StatusOK, resp.Code)

// 	// You can add more assertions to validate the response body
// 	// based on the expected data from your mock database.
// 	// Example: assert.Contains(t, resp.Body.String(), "expectedData")

// 	// Cleanup (if you have any resources to release)
// }

// func TestGetLocationById(t *testing.T) {
// 	// Initialize a Gin router and set its mode to TestMode
// 	r := gin.New()
// 	r.GET("/locations/:location_id", GetLocationById)

// 	// Create a mock database and replace the actual DB with it
// 	// You may want to use a package like github.com/DATA-DOG/go-sqlmock
// 	// to create a mock database for testing.
// 	// Replace models.DB with your mock database instance.

// 	// Perform a GET request to the /locations/:location_id endpoint
// 	req, err := http.NewRequest("GET", "/locations/1", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	resp := httptest.NewRecorder()
// 	r.ServeHTTP(resp, req)

// 	// Check the response status code and body
// 	assert.Equal(t, http.StatusOK, resp.Code)

// 	// You can add more assertions to validate the response body
// 	// based on the expected data from your mock database.
// 	// Example: assert.Contains(t, resp.Body.String(), "expectedData")

// 	// Cleanup (if you have any resources to release)
// }
