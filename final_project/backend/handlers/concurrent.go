package handlers

import (
	"net/http"
	"sync"
	"backend/db"
	"encoding/json"
)

func FetchConcurrentData(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	var users []db.User
	var products []db.Product
	var orders []db.Order
	var fetchError error

	wg.Add(3)

	// Fetch Users
	go func() {
		defer wg.Done()
		if err := db.DB.Find(&users).Error; err != nil {
			fetchError = err
		}
	}()

	// Fetch Products
	go func() {
		defer wg.Done()
		if err := db.DB.Find(&products).Error; err != nil {
			fetchError = err
		}
	}()

	// Fetch Orders
	go func() {
		defer wg.Done()
		if err := db.DB.Find(&orders).Error; err != nil {
			fetchError = err
		}
	}()

	wg.Wait()

	if fetchError != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"users":    users,
		"products": products,
		"orders":   orders,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
