package utils

import (
	"backend/models"
	"errors"
	"log"
	"os"
	"sync"
    "expvar"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
    // "github.com/prometheus/client_golang/prometheus"
)

// Example role constants
const (
    AdminRole = "admin"
    UserRole  = "user"
    GuestRole = "guest"
)

var (
    Users         = make(map[string]string) // Store users with hashed passwords
    Validate      = validator.New()          // Validator instance
    JwtKey        = []byte("your_secret_key") // Secret key for JWT signing
    TokenBlacklist = make(map[string]struct{}) // Blacklist for revoked tokens
    Mu            sync.Mutex // Mutex for concurrent access
    isDevelopment = os.Getenv("ENV") == "development" // Set ENV=development in your dev environment
    Log *logrus.Logger

    // Define expvar variables for metrics
    RequestCount   = expvar.NewInt("request_count")   // Total number of requests
    ResponseTime   = expvar.NewFloat("response_time") // Total response time
    ErrorCount     = expvar.NewInt("error_count")     // Total number of errors
)

func init() {
    // Register metrics with Prometheus
	// prometheus.MustRegister(RequestCount)
	// prometheus.MustRegister(RequestDuration)

    // Initialize Logrus logger
    Log = logrus.New()
    Log.SetFormatter(&logrus.JSONFormatter{}) // Use JSONFormatter for structured logging
    Log.SetOutput(os.Stdout) // Log to standard output

    // Set log level based on environment
    if isDevelopment {
        Log.SetLevel(logrus.DebugLevel)
        Log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
    } else {
        Log.SetLevel(logrus.InfoLevel)
        Log.SetFormatter(&logrus.JSONFormatter{})
    }

    Log.WithFields(logrus.Fields{
        "event": "init_logger",
    }).Info("Logger initialized successfully")
}

// HashPassword hashes the user's password securely
func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        Log.WithFields(logrus.Fields{
            "event": "hash_password",
            "error": err,
        }).Error("Failed to hash password")
        return "", err
    }
    Log.WithFields(logrus.Fields{"event": "hash_password"}).Info("Password hashed successfully")
    return string(hashedPassword), nil
}

// VerifyPassword compares the provided password with the hashed password
func VerifyPassword(hashedPassword, password string) error {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err != nil {
        Log.WithFields(logrus.Fields{
            "event": "verify_password",
            "error": err,
        }).Warn("Password verification failed")
    }
    return err
}

// AddTokenToBlacklist adds a token to the blacklist
func AddTokenToBlacklist(tokenStr string) {
    Mu.Lock()
    defer Mu.Unlock()
    TokenBlacklist[tokenStr] = struct{}{}
    Log.WithFields(logrus.Fields{
        "event": "token_blacklist",
        "token": tokenStr,
    }).Info("Token added to blacklist")
}

// Check if a token is blacklisted
func IsTokenBlacklisted(tokenStr string) bool {
    Mu.Lock()
    defer Mu.Unlock()
    _, exists := TokenBlacklist[tokenStr]
    if exists {
        Log.WithFields(logrus.Fields{
            "event": "check_blacklist",
            "token": tokenStr,
        }).Warn("Token is blacklisted")
    }
    return exists
}

// Extract user role from token
func ExtractUserRoleFromToken(tokenStr string) (string, error) {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return JwtKey, nil
    })
    if err != nil || !token.Valid {
        Log.WithFields(logrus.Fields{
            "event": "extract_user_role",
            "error": err,
        }).Error("Failed to parse token or token invalid")
        return "", errors.New("invalid token")
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        if role, exists := claims["role"]; exists {
            return role.(string), nil
        }
    }
    return "", errors.New("role not found in token")
}

// LogError logs an error message; optionally logs detailed stack trace in development mode
func LogError(msg string, err interface{}) {
    if isDevelopment {
        log.Printf("[ERROR] %s: %+v", msg, err) // Detailed logging for development
    } else {
        log.Printf("[ERROR] %s", msg) // Generic error log for production
    }
}

func GetUsernameFromToken(tokenStr string) (string, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
        return JwtKey, nil
    })
    if err != nil {
        return "", err
    }

    if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
        return claims.Username, nil
    }

    return "", errors.New("invalid token")
}