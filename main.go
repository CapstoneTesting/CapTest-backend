package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type YourStruct struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func main() {
	//Returns a new Mux object that implements the Router interface.
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		http.ServeFile(w, r, "/Users/h.lim.10/Desktop/CapTest/frontend")
	})

	r.Get("/hello", helloHandler)

	// Define the handler for the POST endpoint at "/test".
	r.Post("/test", func(w http.ResponseWriter, r *http.Request) {
		// Parse the JSON request body into a struct.
		var data YourStruct
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Process the data or perform some action.
		// In this example, we'll just send back the received data as a JSON response.
		responseData := map[string]interface{}{
			"receivedData": data,
			"status":       "success",
		}

		// Serialize the response data as JSON and write it to the response.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(responseData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8080", r)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello from Go Backend!"))
}
