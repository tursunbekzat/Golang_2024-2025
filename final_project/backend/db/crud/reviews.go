package crud

import (
	"fmt"
	"backend/db"
)

// Add Review
func AddReview(productID, userID uint, rating int, comment string) {
	review := db.Review{ProductID: productID, UserID: userID, Rating: rating, Comment: comment}
	result := db.DB.Create(&review)
	if result.Error != nil {
		fmt.Println("Error adding review:", result.Error)
		return
	}
	fmt.Println("Review added successfully:", review)
}

// Get All Reviews for Product
func GetProductReviews(productID uint) {
	var reviews []db.Review
	result := db.DB.Find(&reviews, "product_id = ?", productID)
	if result.Error != nil {
		fmt.Println("Error fetching reviews:", result.Error)
		return
	}
	fmt.Println("Product Reviews:", reviews)
}
