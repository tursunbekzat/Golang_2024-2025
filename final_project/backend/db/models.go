package db

import "time"

// Users
type User struct {
	UserID       uint      `gorm:"primaryKey" json:"user_id"`
	Username     string    `gorm:"unique;not null" json:"username"`
	PasswordHash string    `gorm:"not null" json:"password_hash"`
	Email        string    `gorm:"unique;not null" json:"email"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	Role         string    `gorm:"default:customer" json:"role"`
}

// Products
type Product struct {
	ProductID   uint      `gorm:"primaryKey" json:"product_id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	Price       float64   `gorm:"not null" json:"price"`
	Stock       int       `gorm:"default:0" json:"stock"`
	CategoryID  uint      `gorm:"not null;foreignKey:CategoryID;constraint:OnDelete:CASCADE;" json:"category_id"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// Categories
type Category struct {
	CategoryID  uint   `gorm:"primaryKey" json:"category_id"`
	Name        string `gorm:"not null" json:"name"`
	Description string `json:"description"`
}

// Orders
type Order struct {
	OrderID     uint      `gorm:"primaryKey" json:"order_id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	OrderDate   time.Time `gorm:"autoCreateTime" json:"order_date"`
	Status      string    `gorm:"default:pending" json:"status"`
	TotalAmount float64   `gorm:"not null" json:"total_amount"`
}

// Order Items
type OrderItem struct {
	OrderItemID uint    `gorm:"primaryKey" json:"order_item_id"`
	OrderID     uint    `gorm:"not null" json:"order_id"`
	ProductID   uint    `gorm:"not null" json:"product_id"`
	Quantity    int     `gorm:"not null" json:"quantity"`
	Price       float64 `gorm:"not null" json:"price"`
}

// Shopping Cart
type ShoppingCart struct {
	CartID    uint      `gorm:"primaryKey" json:"cart_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// Cart Items
type CartItem struct {
	CartItemID uint `gorm:"primaryKey" json:"cart_item_id"`
	CartID     uint `gorm:"not null" json:"cart_id"`
	ProductID  uint `gorm:"not null" json:"product_id"`
	Quantity   int  `gorm:"not null" json:"quantity"`
}

// Payments
type Payment struct {
	PaymentID     uint      `gorm:"primaryKey" json:"payment_id"`
	OrderID       uint      `gorm:"not null" json:"order_id"`
	Amount        float64   `gorm:"not null" json:"amount"`
	PaymentDate   time.Time `gorm:"autoCreateTime" json:"payment_date"`
	PaymentMethod string    `json:"payment_method"`
}

// Audit Logs
type AuditLog struct {
	LogID     uint      `gorm:"primaryKey" json:"log_id"`
	Action    string    `gorm:"not null" json:"action"`
	UserID    uint      `gorm:"index" json:"user_id"`
	Timestamp time.Time `gorm:"autoCreateTime" json:"timestamp"`
}

// Cache
type Cache struct {
	CacheKey       string    `gorm:"primaryKey" json:"cache_key"`
	CacheValue     string    `gorm:"not null" json:"cache_value"`
	ExpirationTime time.Time `json:"expiration_time"`
}

// Product Images
type ProductImage struct {
	ImageID   uint      `gorm:"primaryKey" json:"image_id"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	ImageURL  string    `gorm:"not null" json:"image_url"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// Reviews
type Review struct {
	ReviewID  uint      `gorm:"primaryKey" json:"review_id"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Rating    int       `gorm:"not null;check:rating>=1 AND rating<=5" json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// Roles
type Role struct {
	RoleID   uint   `gorm:"primaryKey" json:"role_id"`
	RoleName string `gorm:"unique;not null" json:"role_name"`
}

// Sessions
type Session struct {
	SessionID uint      `gorm:"primaryKey" json:"session_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
}

// User Addresses
type UserAddress struct {
	AddressID uint   `gorm:"primaryKey" json:"address_id"`
	UserID    uint   `gorm:"not null" json:"user_id"`
	Street    string `gorm:"not null" json:"street"`
	City      string `gorm:"not null" json:"city"`
	State     string `json:"state"`
	ZipCode   string `gorm:"not null" json:"zip_code"`
}
