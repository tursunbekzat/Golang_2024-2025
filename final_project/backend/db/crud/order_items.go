package crud

import (
	"fmt"
	"backend/db"
)

// Add Item to Order
func AddItemToOrder(orderID, productID uint, quantity int, price float64) {
	item := db.OrderItem{OrderID: orderID, ProductID: productID, Quantity: quantity, Price: price}
	result := db.DB.Create(&item)
	if result.Error != nil {
		fmt.Println("Error adding item to order:", result.Error)
		return
	}
	fmt.Println("Item added to order successfully:", item)
}

// Get All Items in Order
func GetOrderItems(orderID uint) {
	var items []db.OrderItem
	result := db.DB.Find(&items, "order_id = ?", orderID)
	if result.Error != nil {
		fmt.Println("Error fetching order items:", result.Error)
		return
	}
	fmt.Println("Order Items:", items)
}
