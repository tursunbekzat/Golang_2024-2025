package crud

import (
	"fmt"
	"backend/db"
)

// Create Product
func CreateProduct(name string, description string, price float64, stock int, categoryID uint) {
	product := db.Product{Name: name, Description: description, Price: price, Stock: stock, CategoryID: categoryID}
	result := db.DB.Create(&product)
	if result.Error != nil {
		fmt.Println("Error creating product:", result.Error)
		return
	}
	fmt.Println("Product created successfully:", product)
}

// Get All Products
func GetAllProducts() {
	var products []db.Product
	result := db.DB.Find(&products)
	if result.Error != nil {
		fmt.Println("Error fetching products:", result.Error)
		return
	}
	fmt.Println("Products:", products)
}

// Update Product Stock
func UpdateProductStock(productID uint, newStock int) {
	var product db.Product
	result := db.DB.First(&product, productID)
	if result.Error != nil {
		fmt.Println("Product not found:", result.Error)
		return
	}

	product.Stock = newStock
	db.DB.Save(&product)
	fmt.Println("Product stock updated successfully:", product)
}

// Delete Product
func DeleteProduct(productID uint) {
	result := db.DB.Delete(&db.Product{}, productID)
	if result.Error != nil {
		fmt.Println("Error deleting product:", result.Error)
		return
	}
	fmt.Println("Product deleted successfully.")
}
