package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"backend/db"
	"github.com/gorilla/mux"
)

// Get All Order Items
func GetOrderItems(w http.ResponseWriter, r *http.Request) {
	var orderItems []db.OrderItem
	db.DB.Find(&orderItems)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orderItems)
}

// Get Order Items by Order ID
func GetOrderItemsByOrderID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	orderID, _ := strconv.Atoi(params["order_id"])

	var orderItems []db.OrderItem
	db.DB.Where("order_id = ?", orderID).Find(&orderItems)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orderItems)
}

// Create Order Item
func CreateOrderItem(w http.ResponseWriter, r *http.Request) {
	var orderItem db.OrderItem
	json.NewDecoder(r.Body).Decode(&orderItem)

	result := db.DB.Create(&orderItem)
	if result.Error != nil {
		http.Error(w, "Error creating order item", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orderItem)
}

// Delete Order Item
func DeleteOrderItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	result := db.DB.Delete(&db.OrderItem{}, id)
	if result.Error != nil {
		http.Error(w, "Error deleting order item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
