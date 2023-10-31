package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type RequestData struct {
	Message string `json:"message"`
}

func main() {
	//Returns a new Mux object that implements the Router interface.
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Set up CORS middleware with your desired configuration
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // You can specify specific origins here
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300, // Maximum value for MaxAge
	})

	// Use the CORS middleware with your router
	r.Use(corsMiddleware.Handler)

	r.Post("/echo", func(w http.ResponseWriter, r *http.Request) {
		// Create an instance of the RequestData struct to store the JSON data
		var requestData RequestData

		// Parse the JSON data from the request body
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&requestData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Print and echo the received JSON data
		fmt.Printf("Received POST request with JSON data: %s\n", requestData.Message)

		// Echo the JSON data in the response
		responseData := RequestData{Message: requestData.Message}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(responseData)
	})

	// Start the server
	port := ":8200"
	fmt.Printf("Server is listening on port %s\n", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal(err)
	}
}
