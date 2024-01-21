package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"rpg/internal/packcalculator/handlers"
	"rpg/internal/packcalculator/models"
)

// TestCalculateHandler_ValidRequest tests the handling of a valid request.
func TestCalculateHandler_ValidRequest(t *testing.T) {
	// Create a test HTTP request.
	requestBody := `{"order": 263, "pack_sizes": [23, 31, 53]}`
	req, err := http.NewRequest("POST", "/calculate", bytes.NewBufferString(requestBody))
	assert.NoError(t, err)

	// Create a fake HTTP response.
	w := httptest.NewRecorder()

	// Call CalculateHandler.
	handlers.CalculateHandler(w, req)

	// Verify that the response status code is 200 OK.
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the JSON response and check its structure.
	var response models.CalculateResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Check the total quantity calculated in the response.
	var totalQuantity int
	for _, pack := range response.Packs {
		totalQuantity += pack.PackSize * pack.Quantity
	}
	assert.Equal(t, 263, totalQuantity, "unexpected total quantity")

	// Define the expected packs based on the total quantity.
	expectedPacks := []struct {
		PackSize int `json:"pack_size"`
		Quantity int `json:"quantity"`
	}{
		{PackSize: 23, Quantity: 2},
		{PackSize: 31, Quantity: 7},
	}

	// Check that the response JSON includes the correct number of packs.
	assert.Len(t, response.Packs, len(expectedPacks), "unexpected number of pack sizes in the response")

	// Check the values of each pack size and quantity in the response.
	for i, expected := range expectedPacks {
		assert.Equal(t, expected.PackSize, response.Packs[i].PackSize, "unexpected pack size")
		assert.Equal(t, expected.Quantity, response.Packs[i].Quantity, "unexpected quantity")
	}
}

// TestCalculateHandler_InvalidMethod tests the handling of an invalid request method.
func TestCalculateHandler_InvalidMethod(t *testing.T) {
	// Create a test HTTP request with an incorrect method.
	req, err := http.NewRequest("GET", "/calculate", nil)
	assert.NoError(t, err)

	// Create a fake HTTP response.
	w := httptest.NewRecorder()

	// Call CalculateHandler.
	handlers.CalculateHandler(w, req)

	// Verify that the response status code is 405 Method Not Allowed.
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
}

// TestCalculateHandler_InvalidJSON tests the handling of a request with invalid JSON.
func TestCalculateHandler_InvalidJSON(t *testing.T) {
	// Create a test HTTP request with invalid JSON.
	req, err := http.NewRequest("POST", "/calculate", bytes.NewBufferString("invalid json"))
	assert.NoError(t, err)

	// Create a fake HTTP response.
	w := httptest.NewRecorder()

	// Call CalculateHandler.
	handlers.CalculateHandler(w, req)

	// Verify that the response status code is 400 Bad Request.
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
