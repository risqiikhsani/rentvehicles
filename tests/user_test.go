// package tests

// import (
// 	"bytes"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/risqiikhsani/rentvehicles/routes"
// 	"github.com/stretchr/testify/assert"
// )

// func setupTestRouter() *gin.Engine {
// 	r := gin.Default()

// 	// Simulate your route setup
// 	api := r.Group("/api")
// 	routes.SetupUserRoutes(api)

// 	return r
// }

// func TestGetUsers(t *testing.T) {
// 	router := setupTestRouter()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/api/users", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	// You can add more assertions to check the response body
// }

// func TestGetUserById(t *testing.T) {
// 	router := setupTestRouter()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/api/users/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	// You can add more assertions to check the response body
// }

// func TestUpdateUserById(t *testing.T) {
// 	router := setupTestRouter()

// 	// Create a JSON payload to update the user
// 	jsonPayload := []byte(`{"Name": "Updated Name", "Description": "Updated Description"}`)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("PUT", "/api/users/1", bytes.NewBuffer(jsonPayload))
// 	req.Header.Set("Content-Type", "application/json")
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	// You can add more assertions to check the response body and database state
// }

// func TestDeleteUserById(t *testing.T) {
// 	router := setupTestRouter()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/api/users/1", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 204, w.Code)
// 	// You can add more assertions to verify the database state
// }

// func TestCreateUser(t *testing.T) {
// 	router := setupTestRouter()

// 	// Create a JSON payload to create a user
// 	jsonPayload := []byte(`{"Name": "New User", "Description": "User Description"}`)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonPayload))
// 	req.Header.Set("Content-Type", "application/json")
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 201, w.Code)
// 	// You can add more assertions to check the response body and verify the database state
// }
