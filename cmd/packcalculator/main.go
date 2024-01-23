package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"rpg/internal/packcalculator/handlers"
)

func main() {
	// Create a new router from the "gorilla/mux" package.
	router := mux.NewRouter()

	// Handle requests to the '/calculate' endpoint using the CalculateHandler function.
	router.HandleFunc("/calculate", handlers.CalculateHandler).Methods("POST")

	// Use the cors.Default() function to enable CORS with default options.
	corsHandler := cors.Default().Handler(router)

	// Get the port from the environment variable or use a default value (8080).
	port := os.Getenv("RPG_BACKEND_PORT")
	if port == "" {
		port = "8080"
	}

	// Log the information about the server starting.
	log.Printf("Server starting on port %s...\n", port)

	// Start the HTTP server on the specified port with CORS handling.
	if err := http.ListenAndServe(":"+port, corsHandler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
