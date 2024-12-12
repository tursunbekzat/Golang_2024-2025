package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"backend/db"
	"github.com/gorilla/mux"
)

// Get All Payments
func GetPayments(w http.ResponseWriter, r *http.Request) {
	var payments []db.Payment
	db.DB.Find(&payments)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)
}

// Get Payment by ID
func GetPaymentByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var payment db.Payment
	result := db.DB.First(&payment, id)
	if result.Error != nil {
		http.Error(w, "Payment not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payment)
}

// Create Payment
func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment db.Payment
	json.NewDecoder(r.Body).Decode(&payment)

	result := db.DB.Create(&payment)
	if result.Error != nil {
		http.Error(w, "Error creating payment", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payment)
}

// Delete Payment
func DeletePayment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	result := db.DB.Delete(&db.Payment{}, id)
	if result.Error != nil {
		http.Error(w, "Error deleting payment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
