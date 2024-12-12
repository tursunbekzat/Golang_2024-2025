package crud

import (
	"fmt"
	"backend/db"
)

// Add User Address
func AddUserAddress(userID uint, street, city, state, zipCode string) {
	address := db.UserAddress{UserID: userID, Street: street, City: city, State: state, ZipCode: zipCode}
	result := db.DB.Create(&address)
	if result.Error != nil {
		fmt.Println("Error adding user address:", result.Error)
		return
	}
	fmt.Println("User address added successfully:", address)
}

// Get All Addresses for User
func GetUserAddresses(userID uint) {
	var addresses []db.UserAddress
	result := db.DB.Find(&addresses, "user_id = ?", userID)
	if result.Error != nil {
		fmt.Println("Error fetching user addresses:", result.Error)
		return
	}
	fmt.Println("User Addresses:", addresses)
}
