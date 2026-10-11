package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Create a student object
	student := Student{
		ID:   1,
		Name: "Ali",
	}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode the student object into JSON and send it
	if err := json.NewEncoder(w).Encode(student); err != nil {
		log.Printf("Failed to encode student: %v", err)
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
