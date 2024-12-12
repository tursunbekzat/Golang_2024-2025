package crud

import (
	"fmt"
	"backend/db"
)

// Add Item to Cart
func AddItemToCart(cartID, productID uint, quantity int) {
	item := db.CartItem{CartID: cartID, ProductID: productID, Quantity: quantity}
	result := db.DB.Create(&item)
	if result.Error != nil {
		fmt.Println("Error adding item to cart:", result.Error)
		return
	}
	fmt.Println("Item added to cart successfully:", item)
}

// Get All Items in Cart
func GetCartItems(cartID uint) {
	var items []db.CartItem
	result := db.DB.Find(&items, "cart_id = ?", cartID)
	if result.Error != nil {
		fmt.Println("Error fetching cart items:", result.Error)
		return
	}
	fmt.Println("Cart Items:", items)
}

// Remove Item from Cart
func RemoveCartItem(cartItemID uint) {
	result := db.DB.Delete(&db.CartItem{}, cartItemID)
	if result.Error != nil {
		fmt.Println("Error removing cart item:", result.Error)
		return
	}
	fmt.Println("Cart item removed successfully.")
}
