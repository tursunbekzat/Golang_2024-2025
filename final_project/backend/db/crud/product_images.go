package crud

import (
	"fmt"
	"backend/db"
)

// Add Product Image
func AddProductImage(productID uint, imageURL string) {
	image := db.ProductImage{ProductID: productID, ImageURL: imageURL}
	result := db.DB.Create(&image)
	if result.Error != nil {
		fmt.Println("Error adding product image:", result.Error)
		return
	}
	fmt.Println("Product image added successfully:", image)
}

// Get All Images for Product
func GetProductImages(productID uint) {
	var images []db.ProductImage
	result := db.DB.Find(&images, "product_id = ?", productID)
	if result.Error != nil {
		fmt.Println("Error fetching product images:", result.Error)
		return
	}
	fmt.Println("Product Images:", images)
}
