package crud

import (
	"fmt"
	"backend/db"
)

// Create Category
func CreateCategory(name string, description string) {
	category := db.Category{Name: name, Description: description}
	result := db.DB.Create(&category)
	if result.Error != nil {
		fmt.Println("Error creating category:", result.Error)
		return
	}
	fmt.Println("Category created successfully:", category)
}

// Get All Categories
func GetAllCategories() {
	var categories []db.Category
	result := db.DB.Find(&categories)
	if result.Error != nil {
		fmt.Println("Error fetching categories:", result.Error)
		return
	}
	fmt.Println("Categories:", categories)
}

// Update Category
func UpdateCategory(categoryID uint, newName string, newDescription string) {
	var category db.Category
	result := db.DB.First(&category, categoryID)
	if result.Error != nil {
		fmt.Println("Category not found:", result.Error)
		return
	}

	category.Name = newName
	category.Description = newDescription
	db.DB.Save(&category)
	fmt.Println("Category updated successfully:", category)
}

// Delete Category
func DeleteCategory(categoryID uint) {
	result := db.DB.Delete(&db.Category{}, categoryID)
	if result.Error != nil {
		fmt.Println("Error deleting category:", result.Error)
		return
	}
	fmt.Println("Category deleted successfully.")
}
