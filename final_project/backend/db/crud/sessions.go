package crud

import (
	"fmt"
	"backend/db"
	"time"
)

// Create Session
func CreateSession(userID uint, expiresAt time.Time) {
	session := db.Session{UserID: userID, ExpiresAt: expiresAt}
	result := db.DB.Create(&session)
	if result.Error != nil {
		fmt.Println("Error creating session:", result.Error)
		return
	}
	fmt.Println("Session created successfully:", session)
}

// Get All Sessions
func GetAllSessions() {
	var sessions []db.Session
	result := db.DB.Find(&sessions)
	if result.Error != nil {
		fmt.Println("Error fetching sessions:", result.Error)
		return
	}
	fmt.Println("Sessions:", sessions)
}
