package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"Onboarding_Service/routes"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get server port from .env
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// Initialize the router
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	log.Printf("Server is running on port :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
