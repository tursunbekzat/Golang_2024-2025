package main

import (
	"fmt"
	"errors"
)

func calculateTotal(price float64, quantity int) float64 {
    return price * float64(quantity)
}


func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
	//  "Hello, World!" Program
	fmt.Println("Hello, World!")

	// Variables
    var name string = "E-Commerce Platform"
    var version int = 1
    var price float64 = 19.99
    var active bool = true

    fmt.Printf("Name: %s, Version: %d, Price: $%.2f, Active: %t\n", name, version, price, active)

	// Conditional Statements
    stock := 20
    if stock > 0 {
        fmt.Println("Product in stock")
    } else {
        fmt.Println("Out of stock")
    }	

	// Loops
	for i := 1; i <= 5; i++ {
        fmt.Printf("Item %d\n", i)
    }

	// Switch Statements
	category := "Electronics"
    switch category {
    case "Electronics":
        fmt.Println("Category: Electronics")
    case "Clothing":
        fmt.Println("Category: Clothing")
    default:
        fmt.Println("Unknown Category")
    }

	// Functions
    total := calculateTotal(19.99, 3)
    fmt.Printf("Total Price: $%.2f\n", total)

	// Structs
	type Product struct {
        Name  string
        Price float64
        Stock int
    }

	// Error Handling
	result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}