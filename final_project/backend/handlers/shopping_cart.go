package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"backend/db"
	"github.com/gorilla/mux"
)

// Get All Shopping Carts
func GetShoppingCarts(w http.ResponseWriter, r *http.Request) {
	var carts []db.ShoppingCart
	db.DB.Find(&carts)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(carts)
}

// Get Shopping Cart by ID
func GetShoppingCartByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var cart db.ShoppingCart
	result := db.DB.First(&cart, id)
	if result.Error != nil {
		http.Error(w, "Shopping cart not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

// Create Shopping Cart
func CreateShoppingCart(w http.ResponseWriter, r *http.Request) {
	var cart db.ShoppingCart
	json.NewDecoder(r.Body).Decode(&cart)

	result := db.DB.Create(&cart)
	if result.Error != nil {
		http.Error(w, "Error creating shopping cart", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

// Delete Shopping Cart
func DeleteShoppingCart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	result := db.DB.Delete(&db.ShoppingCart{}, id)
	if result.Error != nil {
		http.Error(w, "Error deleting shopping cart", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
