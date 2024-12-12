package main

import "fmt"

// Order struct
type Order struct {
    OrderID   int
    TotalCost float64
}

// PaymentProcessor interface
type PaymentProcessor interface {
    ProcessPayment(amount float64) string
}

// CreditCard implementation
type CreditCard struct {
    CardNumber string
}

func (cc CreditCard) ProcessPayment(amount float64) string {
    return fmt.Sprintf("Order paid using Credit Card ending in %s", cc.CardNumber[len(cc.CardNumber)-4:])
}

// PayPal implementation
type PayPal struct {
    Email string
}

func (pp PayPal) ProcessPayment(amount float64) string {
    return fmt.Sprintf("Order paid using PayPal account %s", pp.Email)
}

// PlaceOrder function
func PlaceOrder(order Order, processor PaymentProcessor) {
    fmt.Printf("Order ID: %d\n", order.OrderID)
    fmt.Println(processor.ProcessPayment(order.TotalCost))
}

func main() {
    order := Order{OrderID: 101, TotalCost: 250.00}

    creditCard := CreditCard{CardNumber: "1234567890123456"}
    PlaceOrder(order, creditCard)

    payPal := PayPal{Email: "user@example.com"}
    PlaceOrder(order, payPal)
}
