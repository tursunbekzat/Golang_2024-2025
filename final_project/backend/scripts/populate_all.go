package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"backend/db"
)

var ctx = context.Background()

// Structs for Mock Data
type MockUser struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
}

type MockProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CategoryID  int     `json:"category_id"`
}

type MockCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MockOrder struct {
	UserID      int    `json:"user_id"`
	OrderDate   string `json:"order_date"`
	Status      string `json:"status"`
	TotalAmount float64 `json:"total_amount"`
}

type MockOrderItem struct {
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type MockPayment struct {
	OrderID       int     `json:"order_id"`
	Amount        float64 `json:"amount"`
	PaymentDate   string  `json:"payment_date"`
	PaymentMethod string  `json:"payment_method"`
}

// Fetch mock data from Mockaroo
// func fetchMockData[T any](url string, result *[]T) error {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return fmt.Errorf("failed to fetch data from Mockaroo: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return fmt.Errorf("failed to read response body: %w", err)
// 	}

// 	if err := json.Unmarshal(body, result); err != nil {
// 		return fmt.Errorf("failed to unmarshal JSON: %w", err)
// 	}

// 	return nil
// }



func fetchMockDataCSV[T any](url string, result *[]T, parser func([]string) T) error {
    resp, err := http.Get(url)
    if err != nil {
        return fmt.Errorf("failed to fetch data from Mockaroo: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := ioutil.ReadAll(resp.Body)
        return fmt.Errorf("unexpected HTTP status: %d, body: %s", resp.StatusCode, string(body))
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return fmt.Errorf("failed to read response body: %w", err)
    }

    reader := csv.NewReader(strings.NewReader(string(body)))
    records, err := reader.ReadAll()
    if err != nil {
        return fmt.Errorf("failed to parse CSV: %w", err)
    }

    for i, record := range records {
        if i == 0 {
            // Skip the header row
            continue
        }
        *result = append(*result, parser(record))
    }

    return nil
}


// Insert Mock Data Functions
func insertUsers(users []MockUser) {
	for _, user := range users {
		newUser := db.User{
			Username:     user.Username,
			Email:        user.Email,
			PasswordHash: user.PasswordHash,
			Role:         user.Role,
		}
		if err := db.DB.Create(&newUser).Error; err != nil {
			log.Printf("Failed to insert user: %v", err)
		} else {
			log.Printf("Inserted user: %s", user.Username)
		}
	}
}

func insertCategories(categories []MockCategory) {
	for _, category := range categories {
		newCategory := db.Category{
			Name:        category.Name,
			Description: category.Description,
		}
		if err := db.DB.Create(&newCategory).Error; err != nil {
			log.Printf("Failed to insert category: %v", err)
		} else {
			log.Printf("Inserted category: %s", category.Name)
		}
	}
}

func insertProducts(products []MockProduct) {
	for _, product := range products {
		newProduct := db.Product{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			CategoryID:  uint(product.CategoryID),
		}
		if err := db.DB.Create(&newProduct).Error; err != nil {
			log.Printf("Failed to insert product: %v", err)
		} else {
			log.Printf("Inserted product: %s", product.Name)
		}
	}
}

func insertOrders(orders []MockOrder) {
	for _, order := range orders {
		orderDate, err := time.Parse(time.RFC3339, order.OrderDate)
		if err != nil {
			log.Printf("Failed to parse order date: %v", err)
			continue
		}
		newOrder := db.Order{
			UserID:      uint(order.UserID),
			OrderDate:   orderDate,
			Status:      order.Status,
			TotalAmount: order.TotalAmount,
		}
		if err := db.DB.Create(&newOrder).Error; err != nil {
			log.Printf("Failed to insert order: %v", err)
		} else {
			log.Printf("Inserted order for user ID: %d", order.UserID)
		}
	}
}
func insertOrderItems(orderItems []MockOrderItem) {
	for _, item := range orderItems {
		newOrderItem := db.OrderItem{
			OrderID:   uint(item.OrderID),
			ProductID: uint(item.ProductID),
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
		if err := db.DB.Create(&newOrderItem).Error; err != nil {
			log.Printf("Failed to insert order item: %v", err)
		} else {
			log.Printf("Inserted order item for order ID: %d", item.OrderID)
		}
	}
}

func insertPayments(payments []MockPayment) {
	for _, payment := range payments {
		paymentDate, err := time.Parse(time.RFC3339, payment.PaymentDate)
		if err != nil {
			log.Printf("Failed to parse payment date: %v", err)
			continue
		}
		newPayment := db.Payment{
			OrderID:       uint(payment.OrderID),
			Amount:        payment.Amount,
			PaymentDate:   paymentDate,
			PaymentMethod: payment.PaymentMethod,
		}
		if err := db.DB.Create(&newPayment).Error; err != nil {
			log.Printf("Failed to insert payment: %v", err)
		} else {
			log.Printf("Inserted payment for order ID: %d", payment.OrderID)
		}
	}
}



func parseMockUser(record []string) MockUser {
    return MockUser{
        Username:     record[0],
        Email:        record[1],
        PasswordHash: record[2],
        Role:         record[3],
    }
}

func parseMockCategory(record []string) MockCategory {
    return MockCategory{
        Name:        record[0],
        Description: record[1],
    }
}

func parseMockProduct(record []string) MockProduct {
    price, _ := strconv.ParseFloat(record[2], 64)
    stock, _ := strconv.Atoi(record[3])
    categoryID, _ := strconv.Atoi(record[4])
    return MockProduct{
        Name:        record[0],
        Description: record[1],
        Price:       price,
        Stock:       stock,
        CategoryID:  categoryID,
    }
}

func parseMockPayment(record []string) MockPayment {
    orderID, _ := strconv.Atoi(record[0])
    amount, _ := strconv.ParseFloat(record[1], 64)
    return MockPayment{
        OrderID:       orderID,
        Amount:        amount,
        PaymentDate:   record[2],
        PaymentMethod: record[3],
    }
}

func parseMockOrderItem(record []string) MockOrderItem {
    orderID, _ := strconv.Atoi(record[0])
    productID, _ := strconv.Atoi(record[1])
    quantity, _ := strconv.Atoi(record[2])
    price, _ := strconv.ParseFloat(record[3], 64)
    return MockOrderItem{
        OrderID:   orderID,
        ProductID: productID,
        Quantity:  quantity,
        Price:     price,
    }
}

func parseMockOrder(record []string) MockOrder {
    userID, _ := strconv.Atoi(record[0])
    totalAmount, _ := strconv.ParseFloat(record[3], 64)
    return MockOrder{
        UserID:      userID,
        OrderDate:   record[1],
        Status:      record[2],
        TotalAmount: totalAmount,
    }
}


func main() {
	// Connect to the database
	db.ConnectDB()

	// Mockaroo API URLs
	userAPI := "https://my.api.mockaroo.com/users.json?key=2d1b2520"
	categoryAPI := "https://my.api.mockaroo.com/categories.json?key=2d1b2520"
	productAPI := "https://my.api.mockaroo.com/products.json?key=2d1b2520"
	orderAPI := "https://my.api.mockaroo.com/orders.json?key=2d1b2520"
	orderItemAPI := "https://my.api.mockaroo.com/order_items.json?key=2d1b2520"
	paymentAPI := "https://my.api.mockaroo.com/payments.json?key=2d1b2520"

	// Fetch and insert data
	var users []MockUser
	var categories []MockCategory
	var products []MockProduct
	var orders []MockOrder
	var orderItems []MockOrderItem
	var payments []MockPayment


	if err := fetchMockDataCSV(userAPI, &users, parseMockUser); err != nil {
		log.Fatalf("Failed to fetch users: %v", err)
	}
	insertUsers(users)

	if err := fetchMockDataCSV(categoryAPI, &categories, parseMockCategory); err != nil {
		log.Fatalf("Failed to fetch categories: %v", err)
	}
	insertCategories(categories)
	

	if err := fetchMockDataCSV(productAPI, &products, parseMockProduct); err != nil {
		log.Fatalf("Failed to fetch products: %v", err)
	}
	insertProducts(products)

	if err := fetchMockDataCSV(orderAPI, &orders, parseMockOrder); err != nil {
		log.Fatalf("Failed to fetch orders: %v", err)
	}
	insertOrders(orders)

	if err := fetchMockDataCSV(orderItemAPI, &orderItems, parseMockOrderItem); err != nil {
		log.Fatalf("Failed to fetch order items: %v", err)
	}
	insertOrderItems(orderItems)

	if err := fetchMockDataCSV(paymentAPI, &payments, parseMockPayment); err != nil {
		log.Fatalf("Failed to fetch payments: %v", err)
	}
	insertPayments(payments)

	log.Println("Data population complete.")
}
