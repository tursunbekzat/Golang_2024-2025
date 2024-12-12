package routes

import (
	"backend/handlers"
	"backend/middlewares"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterRoutes(router *mux.Router) {
	// Public routes
	router.Handle("/metrics", promhttp.Handler())
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST").Handler(middlewares.RateLimitMiddleware(http.HandlerFunc(handlers.LoginHandler)))
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")


	// Concurrent endpoints
	router.HandleFunc("/api/fetch-concurrent", handlers.FetchConcurrentData).Methods("GET")
	router.HandleFunc("/api/process-orders", handlers.ProcessOrders).Methods("POST")

	// Protected routes with authentication
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middlewares.AuthMiddleware)
	
	// Product endpoints
	api.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	// Cache endpoints
	api.HandleFunc("/products/{id}", handlers.GetProductByIDWithCache).Methods("GET")


	// OrderItem endpoints
	api.HandleFunc("/orders", handlers.GetOrderItems).Methods("GET")
	api.HandleFunc("/orders/{id}", handlers.GetOrderItemsByOrderID).Methods("GET")
	api.HandleFunc("/orders", handlers.CreateOrderItem).Methods("POST")
	api.HandleFunc("/orders/{id}", handlers.DeleteOrderItem).Methods("DELETE")

	// Category routes
	api.HandleFunc("/categories", handlers.GetCategories).Methods("GET")
	api.HandleFunc("/categories/{id}", handlers.GetCategoryByID).Methods("GET")
	api.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")
	api.HandleFunc("/categories/{id}", handlers.UpdateCategory).Methods("PUT")
	api.HandleFunc("/categories/{id}", handlers.DeleteCategory).Methods("DELETE")

	// Shopping Cart routes
	api.HandleFunc("/shopping-carts", handlers.GetShoppingCarts).Methods("GET")
	api.HandleFunc("/shopping-carts/{id}", handlers.GetShoppingCartByID).Methods("GET")
	api.HandleFunc("/shopping-carts", handlers.CreateShoppingCart).Methods("POST")
	api.HandleFunc("/shopping-carts/{id}", handlers.DeleteShoppingCart).Methods("DELETE")

	// Payment routes
	api.HandleFunc("/payments", handlers.GetPayments).Methods("GET")
	api.HandleFunc("/payments/{id}", handlers.GetPaymentByID).Methods("GET")
	api.HandleFunc("/payments", handlers.CreatePayment).Methods("POST")
	api.HandleFunc("/payments/{id}", handlers.DeletePayment).Methods("DELETE")

	// Review routes
	api.HandleFunc("/payments", handlers.GetReviews).Methods("GET")
	api.HandleFunc("/payments", handlers.CreateReview).Methods("POST")
	api.HandleFunc("/payments/{id}", handlers.DeleteReview).Methods("DELETE")

	// Category routes
	api.HandleFunc("/categories", handlers.GetCategories).Methods("GET")
	api.HandleFunc("/categories/{id}", handlers.GetCategoryByID).Methods("GET")
	api.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")
	api.HandleFunc("/categories/{id}", handlers.UpdateCategory).Methods("PUT")
	api.HandleFunc("/categories/{id}", handlers.DeleteCategory).Methods("DELETE")

	// Shopping Cart routes
	api.HandleFunc("/shopping-carts", handlers.GetShoppingCarts).Methods("GET")
	api.HandleFunc("/shopping-carts/{id}", handlers.GetShoppingCartByID).Methods("GET")
	api.HandleFunc("/shopping-carts", handlers.CreateShoppingCart).Methods("POST")
	api.HandleFunc("/shopping-carts/{id}", handlers.DeleteShoppingCart).Methods("DELETE")

	// Shopping Cart routes
	api.HandleFunc("/shopping-carts", handlers.GetSessions).Methods("GET")
	api.HandleFunc("/shopping-carts", handlers.CreateSession).Methods("POST")
	api.HandleFunc("/shopping-carts/{id}", handlers.DeleteSession).Methods("DELETE")


	// Admin-only routes
	admin := api.PathPrefix("/admin").Subrouter()
	admin.Use(middlewares.RoleMiddleware("admin"))
	admin.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	admin.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	admin.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")
	admin.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")

	// User endpoints
	admin.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	admin.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
	admin.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	admin.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
}
