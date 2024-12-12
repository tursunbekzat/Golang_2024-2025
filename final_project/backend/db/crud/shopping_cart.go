package crud

import (
	"fmt"
	"backend/db"
)

// Create Shopping Cart
func CreateShoppingCart(userID uint) {
	cart := db.ShoppingCart{UserID: userID}
	result := db.DB.Create(&cart)
	if result.Error != nil {
		fmt.Println("Error creating shopping cart:", result.Error)
		return
	}
	fmt.Println("Shopping cart created successfully:", cart)
}

// Get All Shopping Carts
func GetAllShoppingCarts() {
	var carts []db.ShoppingCart
	result := db.DB.Find(&carts)
	if result.Error != nil {
		fmt.Println("Error fetching shopping carts:", result.Error)
		return
	}
	fmt.Println("Shopping carts:", carts)
}

// Delete Shopping Cart
func DeleteShoppingCart(cartID uint) {
	result := db.DB.Delete(&db.ShoppingCart{}, cartID)
	if result.Error != nil {
		fmt.Println("Error deleting shopping cart:", result.Error)
		return
	}
	fmt.Println("Shopping cart deleted successfully.")
}
