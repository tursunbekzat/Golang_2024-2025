package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"backend/db"
	"backend/middlewares"
	"backend/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Define Prometheus metrics
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path"},
	)
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

func init() {
	// Configure Logrus
	log.SetFormatter(&log.JSONFormatter{}) // Use JSON format for logs
	log.SetOutput(os.Stdout)              // Log to stdout
	log.SetLevel(log.InfoLevel)           // Set default log level

	// Register Prometheus metrics
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(requestDuration)
}

func main() {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	db.ConnectDB()

	// Connect to Redis
	db.ConnectRedis()

	// Auto-migrate tables
	db.DB.AutoMigrate(&db.User{}, &db.Product{}, &db.Category{}, &db.Order{}, &db.OrderItem{}, &db.ShoppingCart{}, &db.CartItem{}, &db.Payment{})

	// Create a new router
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	// Add security headers middleware
	router.Use(middlewares.SecurityHeadersMiddleware)

	// Add logging middleware
	router.Use(middlewares.RequestLoggingMiddleware)


	// Configure CORS
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),        // Allow frontend origin
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}), // Allowed HTTP methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), // Allowed headers
	)

	// Example route to test metrics
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		requestCounter.WithLabelValues(r.Method, "/test").Inc()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Test endpoint"))
	})

	// Create an HTTP server
	server := &http.Server{
		Addr:    ":8080",
		Handler: corsOptions(router),
	}

	// Graceful shutdown
	go func() {
		log.Println("Server is running on port 8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on port 8080: %v\n", err)
		}
	}()

	// Listen for OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}

// prometheusMiddleware collects metrics for every request
func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)

		// Update Prometheus metrics
		requestCounter.WithLabelValues(r.Method, r.URL.Path).Inc()
		requestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration.Seconds())
	})
}