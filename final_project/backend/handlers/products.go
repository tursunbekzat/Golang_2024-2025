package handlers

import (
	"backend/db"
	
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

)

// Get All Products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []db.Product
	db.DB.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// Get Product by ID
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var product db.Product
	result := db.DB.First(&product, id)
	if result.Error != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// Create Product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product db.Product
	json.NewDecoder(r.Body).Decode(&product)

	result := db.DB.Create(&product)
	if result.Error != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// Update Product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var product db.Product
	result := db.DB.First(&product, id)
	if result.Error != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&product)
	db.DB.Save(&product)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

	// Invalidate cache
	productID := r.URL.Query().Get("id")
	cacheKey := "product:" + productID
	err := db.RedisClient.Del(ctx, cacheKey).Err()
	if err != nil {
		log.Printf("Failed to invalidate cache for product ID %s: %v", productID, err)
	}
}

// Delete Product
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	result := db.DB.Delete(&db.Product{}, id)
	if result.Error != nil {
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}





var ctx = context.Background()

// GetProductByIDWithCache retrieves product details from cache or database
func GetProductByIDWithCache(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")
	if productID == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	// Check Redis cache
	cacheKey := "product:" + productID
	cachedData, err := db.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		// Return cached data
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(cachedData))
		return
	}

	// If not in cache, fetch from database
	var product db.Product
	result := db.DB.First(&product, "product_id = ?", productID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch product", http.StatusInternalServerError)
		}
		return
	}

	// Serialize product to JSON
	productJSON, err := json.Marshal(product)
	if err != nil {
		http.Error(w, "Failed to serialize product", http.StatusInternalServerError)
		return
	}

	// Cache product data
	err = db.RedisClient.Set(ctx, cacheKey, productJSON, 10*time.Minute).Err()
	if err != nil {
		http.Error(w, "Failed to cache product data", http.StatusInternalServerError)
		return
	}

	// Return product data
	w.Header().Set("Content-Type", "application/json")
	w.Write(productJSON)
}