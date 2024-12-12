package crud

import (
	"fmt"
	"backend/db"
)

// Create User
func CreateUser(username, passwordHash, email string) {
	user := db.User{Username: username, PasswordHash: passwordHash, Email: email}
	result := db.DB.Create(&user)
	if result.Error != nil {
		fmt.Println("Error creating user:", result.Error)
		return
	}
	fmt.Println("User created successfully:", user)
}
// Get All Users
func GetAllUsers() {
	var users []db.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		fmt.Println("Error fetching users:", result.Error)
		return
	}
	fmt.Println("Users:", users)
}

// Update User
func UpdateUser(userID uint, newRole string) {
	var user db.User
	result := db.DB.First(&user, userID)
	if result.Error != nil {
		fmt.Println("User not found:", result.Error)
		return
	}

	user.Role = newRole
	db.DB.Save(&user)
	fmt.Println("User updated successfully:", user)
}

// Delete User
func DeleteUser(userID uint) {
	result := db.DB.Delete(&db.User{}, userID)
	if result.Error != nil {
		fmt.Println("Error deleting user:", result.Error)
		return
	}
	fmt.Println("User deleted successfully.")
}