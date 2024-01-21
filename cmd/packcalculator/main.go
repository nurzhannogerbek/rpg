package main

import (
	"log"
	"net/http"

	"rpg/internal/packcalculator/handlers"
)

func main() {
	// Handle requests to the '/calculate' endpoint using the CalculateHandler function.
	http.HandleFunc("/calculate", handlers.CalculateHandler)

	// Start the HTTP server on port 8080.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
