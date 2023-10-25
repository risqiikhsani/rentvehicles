package controllers

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestGetPosts(t *testing.T) {
	// Create a test Gin router
	r := gin.New()
	r.GET("/posts", GetPosts)

	// Create a mock HTTP request
	req := httptest.NewRequest("GET", "/posts", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response status code and body
	require.Equal(t, 200, w.Code)
	// You can add more assertions to check the response body if needed
}

func TestGetPostById(t *testing.T) {
	// Create a test Gin router
	r := gin.New()
	r.GET("/posts/:post_id", GetPostById)

	// Create a mock HTTP request with a post_id parameter
	req := httptest.NewRequest("GET", "/posts/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response status code and body
	require.Equal(t, 200, w.Code)
	// You can add more assertions to check the response body if needed

	// Test case for a non-existent post
	req = httptest.NewRequest("GET", "/posts/999", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response status code for a non-existent post
	require.Equal(t, 404, w.Code)
	// You can check the response body for an error message as well
}
