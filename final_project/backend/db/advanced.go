package db


import (
	"fmt"
)

func GetOrdersWithUserDetails() {
	var orders []Order
	result := DB.Preload("User").Find(&orders) // Preloads related User for each Order
	if result.Error != nil {
		fmt.Println("Error fetching orders:", result.Error)
		return
	}

	for _, order := range orders {
		fmt.Printf("Order ID: %d, User: %d, Total Amount: %.2f, Status: %s\n",
			order.OrderID, order.UserID, order.TotalAmount, order.Status)
	}
}

func CalculateTotalSales(productID uint) {
	var totalSales float64
	result := DB.Model(&OrderItem{}).
		Select("SUM(price * quantity)").
		Where("product_id = ?", productID).
		Scan(&totalSales)
	if result.Error != nil {
		fmt.Println("Error calculating total sales:", result.Error)
		return
	}

	fmt.Printf("Total Sales for Product ID %d: %.2f\n", productID, totalSales)
}

func CountOrdersForUser(userID uint) {
	var orderCount int64
	result := DB.Model(&Order{}).Where("user_id = ?", userID).Count(&orderCount)
	if result.Error != nil {
		fmt.Println("Error counting orders:", result.Error)
		return
	}

	fmt.Printf("User ID %d has placed %d orders\n", userID, orderCount)
}

func GetProductsByCategory(categoryID uint) {
	var products []Product
	result := DB.Where("category_id = ?", categoryID).Find(&products)
	if result.Error != nil {
		fmt.Println("Error fetching products:", result.Error)
		return
	}

	for _, product := range products {
		fmt.Printf("Product: %s, Price: %.2f, Stock: %d\n", product.Name, product.Price, product.Stock)
	}
}

func GetProductsSortedByPrice(order string) {
	var products []Product
	result := DB.Order("price " + order).Find(&products)
	if result.Error != nil {
		fmt.Println("Error fetching sorted products:", result.Error)
		return
	}

	for _, product := range products {
		fmt.Printf("Product: %s, Price: %.2f\n", product.Name, product.Price)
	}
}

func GetPaginatedProducts(page, pageSize int) {
	var products []Product
	offset := (page - 1) * pageSize
	result := DB.Offset(offset).Limit(pageSize).Find(&products)
	if result.Error != nil {
		fmt.Println("Error fetching paginated products:", result.Error)
		return
	}

	for _, product := range products {
		fmt.Printf("Product: %s, Price: %.2f\n", product.Name, product.Price)
	}
}
