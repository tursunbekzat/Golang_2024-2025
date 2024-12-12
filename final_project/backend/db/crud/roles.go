package crud

import (
	"fmt"
	"backend/db"
)

// Add Role
func AddRole(roleName string) {
	role := db.Role{RoleName: roleName}
	result := db.DB.Create(&role)
	if result.Error != nil {
		fmt.Println("Error adding role:", result.Error)
		return
	}
	fmt.Println("Role added successfully:", role)
}

// Get All Roles
func GetAllRoles() {
	var roles []db.Role
	result := db.DB.Find(&roles)
	if result.Error != nil {
		fmt.Println("Error fetching roles:", result.Error)
		return
	}
	fmt.Println("Roles:", roles)
}
