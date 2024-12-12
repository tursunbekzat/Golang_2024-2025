package handlers

import "fmt"

import (
	"net/http"
	"sync"
	"backend/db"
	"encoding/json"
	"time"
)

// Concurrent Order Processing
func ProcessOrders(w http.ResponseWriter, r *http.Request) {
	var orders []db.Order
	var wg sync.WaitGroup

	if err := db.DB.Find(&orders).Error; err != nil {
		http.Error(w, "Failed to fetch orders", http.StatusInternalServerError)
		return
	}

	response := make(chan string, len(orders))

	// Process each order in a Goroutine
	for _, order := range orders {
		wg.Add(1)
		go func(order db.Order) {
			defer wg.Done()
			time.Sleep(2 * time.Second) // Simulate processing time
			response <- "Processed order ID: " + fmt.Sprint(order.OrderID)
		}(order)
	}

	go func() {
		wg.Wait()
		close(response)
	}()

	var results []string
	for res := range response {
		results = append(results, res)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
