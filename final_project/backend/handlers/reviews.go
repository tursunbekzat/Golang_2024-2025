package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"backend/db"
	"github.com/gorilla/mux"
)

// Get All Reviews
func GetReviews(w http.ResponseWriter, r *http.Request) {
	var reviews []db.Review
	db.DB.Find(&reviews)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}

// Create Review
func CreateReview(w http.ResponseWriter, r *http.Request) {
	var review db.Review
	json.NewDecoder(r.Body).Decode(&review)

	result := db.DB.Create(&review)
	if result.Error != nil {
		http.Error(w, "Error creating review", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(review)
}

// Delete Review
func DeleteReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	result := db.DB.Delete(&db.Review{}, id)
	if result.Error != nil {
		http.Error(w, "Error deleting review", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
