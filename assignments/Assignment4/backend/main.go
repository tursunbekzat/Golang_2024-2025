package main

import (
	"backend/handlers"
	"backend/middleware"
	"backend/utils"
	"fmt"
	"net/http"
    "expvar"

	"github.com/go-playground/validator/v10"
	// "github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    // Create a new router
    r := mux.NewRouter()

    // Create a new validator
    utils.Validate = validator.New()

    // CSRF Protection
    // csrfMiddleware := csrf.Protect([]byte("32-byte-long-secret-key"), csrf.Secure(true))
    csrfMiddleware := middleware.CSRFProtection

    // r.Use(csrf.Protect([]byte("32-byte-long-secret"), csrfOptions))

    // Apply middlewares
    r.Use(middleware.ErrorHandlerMiddleware)
    r.Use(middleware.SecurityHeadersMiddleware)
    r.Use(csrfMiddleware)
    r.Use(middleware.LoggingMiddleware)

    // Expose expvar metrics
    r.Handle("/debug/vars", expvar.Handler()).Methods("GET")

    // Define routes
    r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
    r.HandleFunc("/register", handlers.RegisterHandler).Methods("GET", "POST")
    r.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST")
    r.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET") // Logout route
       
    // Apply AuthMiddleware to routes that require authentication
    protected := r.PathPrefix("/protected").Subrouter()
    protected.Use(middleware.AuthMiddleware) // Apply AuthMiddleware here
    protected.Handle("/admin", middleware.RBACMiddleware("admin", http.HandlerFunc(handlers.ProtectedHandler))).Methods("GET")
   

    // Metrics endpoint
    r.Handle("/metrics", promhttp.Handler()).Methods("GET")

    // Start the server
    fmt.Println("Server is listening on port 8443...")
    err := http.ListenAndServeTLS(":8443", "openssl/server.crt", "openssl/server.key", r)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}
