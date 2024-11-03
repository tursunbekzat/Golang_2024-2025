package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	// "strconv"

	// "github.com/gorilla/mux"
	"rest-api/models"
	// "rest-api/utils"
)

func GetUsersSQL(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := "SELECT id, name, age, email FROM users"
		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []models.User
		for rows.Next() {
			var user models.User
			if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email); err != nil {
				http.Error(w, "Error scanning row", http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}
		json.NewEncoder(w).Encode(users)
	}
}

func CreateUserSQL(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		query := "INSERT INTO users (name, age, email) VALUES ($1, $2, $3)"
		_, err := db.Exec(query, user.Name, user.Age, user.Email)
		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "User created successfully")
	}
}
