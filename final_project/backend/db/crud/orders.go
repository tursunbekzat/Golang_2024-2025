package crud

import (
	"fmt"
	"backend/db"
)

// Create Order
func CreateOrder(userID uint, totalAmount float64) {
	order := db.Order{UserID: userID, TotalAmount: totalAmount}
	result := db.DB.Create(&order)
	if result.Error != nil {
		fmt.Println("Error creating order:", result.Error)
		return
	}
	fmt.Println("Order created successfully:", order)
}

// Get All Orders
func GetAllOrders() {
	var orders []db.Order
	result := db.DB.Find(&orders)
	if result.Error != nil {
		fmt.Println("Error fetching orders:", result.Error)
		return
	}
	fmt.Println("Orders:", orders)
}

// Update Order Status
func UpdateOrderStatus(orderID uint, newStatus string) {
	var order db.Order
	result := db.DB.First(&order, orderID)
	if result.Error != nil {
		fmt.Println("Order not found:", result.Error)
		return
	}

	order.Status = newStatus
	db.DB.Save(&order)
	fmt.Println("Order status updated successfully:", order)
}

// Delete Order
func DeleteOrder(orderID uint) {
	result := db.DB.Delete(&db.Order{}, orderID)
	if result.Error != nil {
		fmt.Println("Error deleting order:", result.Error)
		return
	}
	fmt.Println("Order deleted successfully.")
}
