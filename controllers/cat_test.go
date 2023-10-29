package controllers

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"gorm.io/gorm"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/risqiikhsani/rentvehicles/models"
// 	"github.com/risqiikhsani/rentvehicles/handlers"
// 	"github.com/risqiikhsani/rentvehicles/utils"
// )

// // MockedDB is a mocked implementation of the database for testing.
// type MockedDB struct {
// 	mock.Mock
// }

// func (m *MockedDB) Preload(column string, conditions ...interface{}) *gorm.DB {
// 	args := m.Called(column, conditions)
// 	return args.Get(0).(*gorm.DB)
// }

// func (m *MockedDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
// 	args := m.Called(dest, conds)
// 	return args.Get(0).(*gorm.DB)
// }

// func (m *MockedDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
// 	args := m.Called(dest, conds)
// 	return args.Get(0).(*gorm.DB)
// }

// func (m *MockedDB) Create(dest interface{}) *gorm.DB {
// 	args := m.Called(dest)
// 	return args.Get(0).(*gorm.DB)
// }

// func (m *MockedDB) Where(dest interface{}, conds ...interface{}) *gorm.DB {
// 	args := m.Called(dest, conds)
// 	return args.Get(0).(*gorm.DB)
// }

// func (m *MockedDB) Save(dest interface{}) *gorm.DB {
// 	args := m.Called(dest)
// 	return args.Get(0).(*gorm.DB)
// }

// func (m *MockedDB) Delete(dest interface{}, conds ...interface{}) *gorm.DB {
// 	args := m.Called(dest, conds)
// 	return args.Get(0).(*gorm.DB)
// }

// func TestGetCats(t *testing.T) {
// 	// Create a new instance of the Gin engine
// 	r := gin.Default()

// 	// Create a mocked database
// 	mockDB := new(MockedDB)

// 	// Replace the actual database with the mocked database
// 	models.DB = mockDB

// 	// Define the test data
// 	expectedCats := []models.Cat{
// 		// Define your test data here
// 	}

// 	// Set up expectations for the mocked database
// 	mockDB.On("Preload", "Images").Return(mockDB)
// 	mockDB.On("Find", &[]models.Cat{}).Return(&gorm.DB{
// 		Error: nil,
// 		Value: expectedCats,
// 	})

// 	// Perform a GET request to the GetCats endpoint
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/cats", nil)
// 	r.ServeHTTP(w, req)

// 	// Assert that the response status code is 200 (OK)
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	// Add more assertions to check the response body or other details as needed
// 	// For example, you can assert the response body JSON matches the expected data.
// 	// assert.JSONEq(t, expectedResponse, w.Body.String())

// 	// Verify that the expectations for the mocked database were met
// 	mockDB.AssertExpectations(t)
// }
