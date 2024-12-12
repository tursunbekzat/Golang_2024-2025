package handlers

import (
	"encoding/json"
	"net/http"
	
	"backend/db"
	"backend/auth"

	log "github.com/sirupsen/logrus"
)


func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	// Parse the request body
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := auth.HashPassword(input.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Create a new user object
	user := db.User{
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: hashedPassword,
		Role:         input.Role, // Optional: Default to "customer"
	}

	// Save the user to the database
	if err := db.DB.Create(&user).Error; err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

// LoginHandler handles user login and generates a JWT
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

	// Parse the request body
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.WithError(err).Error("Failed to parse login request")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var user db.User
	if err := db.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		log.WithFields(log.Fields{
			"email": input.Email,
		}).Error("User not found during login")
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := auth.ComparePasswords(user.PasswordHash, input.Password); err != nil {
		log.WithFields(log.Fields{
			"email": input.Email,
		}).Warn("Incorrect password during login")
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateJWT(user.UserID, user.Role)
	if err != nil {
		log.WithError(err).Error("Failed to generate JWT")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{
		"user_id": user.UserID,
		"email":   user.Email,
	}).Info("User logged in successfully")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
