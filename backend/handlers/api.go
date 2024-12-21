package handlers

import (
    "fmt"
    "net/http"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, `{"message": "Hello, API!"}`)
}
