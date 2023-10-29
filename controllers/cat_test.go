package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/models"
	"github.com/risqiikhsani/rentvehicles/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockedDB is a mocked implementation of the database for testing.
type MockedDB struct {
	mock.Mock
}

func (m *MockedDB) GetCats() ([]models.Cat, error) {
	args := m.Called()
	return args.Get(0).([]models.Cat), args.Error(1)
}

func (m *MockedDB) GetCatByID(catID string) (*models.Cat, error) {
	args := m.Called(catID)
	return args.Get(0).(*models.Cat), args.Error(1)
}

func (m *MockedDB) CreateCat(cat *models.Cat) error {
	args := m.Called(cat)
	return args.Error(0)
}

func (m *MockedDB) UpdateCat(cat *models.Cat) error {
	args := m.Called(cat)
	return args.Error(0)
}

func (m *MockedDB) DeleteCat(cat *models.Cat) error {
	args := m.Called(cat)
	return args.Error(0)
}

type MockedAuth struct {
	mock.Mock
}

func (m *MockedAuth) RequireAuthentication2(c *gin.Context, requiredRole string) (uint, string, bool) {
	args := m.Called(c, requiredRole)
	return args.Get(0).(uint), args.Get(1).(string), args.Bool(2)
}

func (m *MockedAuth) CheckAuthentication2(c *gin.Context) (uint, string, bool) {
	args := m.Called(c)
	return args.Get(0).(uint), args.Get(1).(string), args.Bool(2)
}

func TestGetCatsController(t *testing.T) {
	// Create a new instance of the mock CatDatabase
	mockDB := new(MockedDB)
	mockAuth := new(MockedAuth)
	// Create the controller and pass the mockDB
	controller := GetCats(mockDB, mockAuth)

	t.Run("Successful GetCats", func(t *testing.T) {
		// Define the expected data to be returned by the mock
		expectedCats := []models.Cat{
			{Name: "Cat1", Age: 3, Text: "Description1"},
			{Name: "Cat2", Age: 2, Text: "Description2"},
		}

		// Mock the GetCats method to return the expected data
		mockDB.On("GetCats").Return(expectedCats, nil)

		// Create a response recorder
		w := httptest.NewRecorder()

		// Set up a Gin context and call the controller
		c, _ := gin.CreateTestContext(w)
		controller(c)

		// Assert the response status code
		assert.Equal(t, http.StatusOK, w.Code)

		// Unmarshal the response body and compare it with the expected data
		var responseCats []models.Cat
		if err := json.NewDecoder(w.Body).Decode(&responseCats); err != nil {
			t.Fatalf("Error decoding response body: %v", err)
		}
		assert.Equal(t, expectedCats, responseCats)

		// Assert that the GetCats method was called
		mockDB.AssertCalled(t, "GetCats")
	})
}

func TestCreateCatHandler(t *testing.T) {
	utils.InitializeTranslator()
	// Initialize the validator
	utils.InitializeValidator()

	// Create a new instance of the mock CatDatabase
	mockDB := new(MockedDB)
	mockAuth := new(MockedAuth)
	// Create the controller and pass the mockDB
	controller := CreateCat(mockDB, mockAuth) // Use CreateCat instead of GetCats

	t.Run("Successful CreateCat", func(t *testing.T) {
		// Create a new cat for testing
		newCat := models.Cat{
			Name:   "TestCat",
			Age:    2,
			Text:   "Test description",
			UserID: 1,
		}

		// Create a request body with the new cat data
		requestBody, _ := json.Marshal(newCat)

		// Create a new HTTP request with the request body
		req, _ := http.NewRequest("POST", "/api/cats", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		// Create a response recorder
		w := httptest.NewRecorder()

		// Set up a Gin context and call the controller
		c, _ := gin.CreateTestContext(w)
		c.Request = req

		// Mock the RequireAuthentication2 function
		mockAuth.On("RequireAuthentication2", c, "").Return(uint(1), "user", true)
		// Mock the CreateCat method to expect a call with the newCat object and return nil (success).
		mockDB.On("CreateCat", &newCat).Return(nil)
		controller(c)

		// Assert the response status code
		assert.Equal(t, http.StatusCreated, w.Code)

		// Unmarshal the response body and compare it with the new cat data
		var responseCat models.Cat
		if err := json.NewDecoder(w.Body).Decode(&responseCat); err != nil {
			t.Fatalf("Error decoding response body: %v", err)
		}
		assert.Equal(t, newCat, responseCat)

		// Assert that the CreateCat method was called
		mockDB.AssertCalled(t, "CreateCat", &newCat)
	})
}
