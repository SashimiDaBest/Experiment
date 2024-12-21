package main

import (
    "fmt"
    "net/http"

	"github.com/gorilla/mux"
	"backend/handlers"
)

func main() {
	// Register routes
	r := mux.NewRouter()
    r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
    r.HandleFunc("/api", handlers.APIHandler).Methods("GET")

	// http.HandleFunc("/", handlers.HomeHandler)
    // http.HandleFunc("/api", handlers.APIHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        fmt.Println("Error starting server:", err)
    }

    // fmt.Println("Server is running on http://localhost:8080")
    // if err := http.ListenAndServe(":8080", nil); err != nil {
    //     fmt.Println("Error starting server:", err)
    // }
}

