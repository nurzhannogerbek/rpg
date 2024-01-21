package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"rpg/internal/packcalculator/services"
)

// CalculateHandler handles the '/calculate' endpoint.
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST, return Method Not Allowed if not.
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method. Use POST.", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON request body into a struct.
	var request struct {
		Order     int   `json:"order"`
		PackSizes []int `json:"pack_sizes"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// If there is an error decoding JSON, return a Bad Request response.
		http.Error(w, "Error decoding JSON request", http.StatusBadRequest)
		return
	}

	// Call the CalculatePacks function to calculate the optimal packing of sizes.
	result, err := services.CalculatePacks(request.Order, request.PackSizes)
	if err != nil {
		// If an error occurs during calculation, return an Internal Server Error response.
		http.Error(w, "Error calculating packs", http.StatusInternalServerError)
		return
	}

	// Convert the result to JSON format.
	response, err := json.Marshal(result)
	if err != nil {
		// If encoding the response to JSON fails, return an Internal Server Error response.
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		log.Println("Error encoding response:", err)
		return
	}

	// Set the Content-Type header to indicate that the response is in JSON format.
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client.
	_, err = w.Write(response)
	if err != nil {
		// If writing the response fails, log the error and return an Internal Server Error response.
		log.Println("Error writing response:", err)
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}
