package crud

import (
	"fmt"
	"backend/db"
)

// Log an action
func LogAction(userID uint, action string) {
	log := db.AuditLog{UserID: userID, Action: action}
	result := db.DB.Create(&log)
	if result.Error != nil {
		fmt.Println("Error logging action:", result.Error)
		return
	}
	fmt.Println("Action logged successfully:", log)
}

// Get All Logs
func GetAllLogs() {
	var logs []db.AuditLog
	result := db.DB.Find(&logs)
	if result.Error != nil {
		fmt.Println("Error fetching logs:", result.Error)
		return
	}
	fmt.Println("Audit Logs:", logs)
}
