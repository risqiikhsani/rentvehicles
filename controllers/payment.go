package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/risqiikhsani/rentvehicles/configs"
	"github.com/risqiikhsani/rentvehicles/handlers"
	"github.com/risqiikhsani/rentvehicles/models"
	"gorm.io/gorm/clause"
)

func CreatePayment(c *gin.Context) {
	userID, _, authenticated := handlers.RequireAuthentication(c, "Basic")
	if !authenticated {
		return
	}

	rent_id := c.Param("rent_id")

	var rent models.Rent

	result := models.DB.Preload(clause.Associations).First(&rent, rent_id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Rent not found"})
		return
	}

	if rent.UserID != userID {
		c.JSON(400, gin.H{"error": "Not authorized"})
		return
	}

	var rent_id_string string = strconv.FormatUint(uint64(rent.ID), 10)

	data, err := SendPaymentRequest(rent_id_string, rent.RentDetail.EstimatedFinalPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send payment request"})
		return
	}

	c.JSON(201, data)

}

func SendPaymentRequest(orderID string, amount uint) (map[string]interface{}, error) {
	secretsPath := "../"
	secretConfig, err := configs.LoadSecretConfig(secretsPath)
	if err != nil {
		return nil, err
	}
	// Get environment variables or provide default values.
	midtransServerKey := secretConfig.MidtransServerKey
	serverKey := midtransServerKey // Replace with your actual server key

	authString := base64.StdEncoding.EncodeToString([]byte(serverKey + ":"))

	// Define the payload
	payload := map[string]interface{}{
		"transaction_details": map[string]interface{}{
			"order_id":     orderID,
			"gross_amount": amount,
		},
		"credit_card": map[string]interface{}{
			"secure": true,
		},
	}

	// Convert payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// Set up the request
	url := "https://app.sandbox.midtrans.com/snap/v1/transactions"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("accept", "application/json")
	req.Header.Set("authorization", "Basic "+authString)
	req.Header.Set("content-type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	var responseBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
