package handlers

import (
	"encoding/json"
	"net/http"
	"gorm.io/gorm"
	"rest-api/models"
)

func GetUsersGORM(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []models.User
		db.Find(&users)
		json.NewEncoder(w).Encode(users)
	}
}

func CreateUserGORM(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		db.Create(&user)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}
