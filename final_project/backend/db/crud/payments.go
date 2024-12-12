package crud

import (
	"fmt"
	"backend/db"
)

// Create Payment
func CreatePayment(orderID uint, amount float64, paymentMethod string) {
	payment := db.Payment{OrderID: orderID, Amount: amount, PaymentMethod: paymentMethod}
	result := db.DB.Create(&payment)
	if result.Error != nil {
		fmt.Println("Error creating payment:", result.Error)
		return
	}
	fmt.Println("Payment created successfully:", payment)
}

// Get All Payments
func GetAllPayments() {
	var payments []db.Payment
	result := db.DB.Find(&payments)
	if result.Error != nil {
		fmt.Println("Error fetching payments:", result.Error)
		return
	}
	fmt.Println("Payments:", payments)
}

// Delete Payment
func DeletePayment(paymentID uint) {
	result := db.DB.Delete(&db.Payment{}, paymentID)
	if result.Error != nil {
		fmt.Println("Error deleting payment:", result.Error)
		return
	}
	fmt.Println("Payment deleted successfully.")
}
