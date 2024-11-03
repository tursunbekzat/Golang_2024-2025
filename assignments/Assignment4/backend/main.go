package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", helloHandler) // Set up the route
    fmt.Println("Server is listening on port 8080...")
    err := http.ListenAndServe(":8080", nil) // Start the server
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}