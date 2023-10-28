package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
		w.Write([]byte("Hello World!"))
	})

	r.Post("/test", func(w http.ResponseWriter, r *http.Request) {
		var data YourStruct
		// Parse the request body
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}

		// Send a response
		w.WriteHeader(http.StatusOK)
		fmt.Println(w, "Result:", data)
	})

	http.ListenAndServe(":8080", r)
}
