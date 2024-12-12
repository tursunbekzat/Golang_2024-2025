package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"backend/db"
	"github.com/gorilla/mux"
)

// Get All Categories
func GetCategories(w http.ResponseWriter, r *http.Request) {
	var categories []db.Category
	db.DB.Find(&categories)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

// Get Category by ID
func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var category db.Category
	result := db.DB.First(&category, id)
	if result.Error != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

// Create Category
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category db.Category
	json.NewDecoder(r.Body).Decode(&category)

	result := db.DB.Create(&category)
	if result.Error != nil {
		http.Error(w, "Error creating category", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

// Update Category
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var category db.Category
	result := db.DB.First(&category, id)
	if result.Error != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&category)
	db.DB.Save(&category)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

// Delete Category
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	result := db.DB.Delete(&db.Category{}, id)
	if result.Error != nil {
		http.Error(w, "Error deleting category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
