package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"backend/db"
	"github.com/gorilla/mux"
)

// Get All Sessions
func GetSessions(w http.ResponseWriter, r *http.Request) {
	var sessions []db.Session
	db.DB.Find(&sessions)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessions)
}

// Create Session
func CreateSession(w http.ResponseWriter, r *http.Request) {
	var session db.Session
	json.NewDecoder(r.Body).Decode(&session)

	result := db.DB.Create(&session)
	if result.Error != nil {
		http.Error(w, "Error creating session", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}

// Delete Session
func DeleteSession(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	result := db.DB.Delete(&db.Session{}, id)
	if result.Error != nil {
		http.Error(w, "Error deleting session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
