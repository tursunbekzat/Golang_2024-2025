package middleware

import (
    "net/http"
    "backend/utils"
    "backend/models"
    "time"
    "strings"

    "github.com/dgrijalva/jwt-go"
    "github.com/sirupsen/logrus"
    "github.com/gorilla/csrf"
)

// Middleware for checking authentication
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var tokenStr string

        if strings.HasPrefix(r.URL.Path, "/debug/vars") {
            next.ServeHTTP(w, r)
            return
        }

        // First, try to get the token from the Authorization header
        authHeader := r.Header.Get("Authorization")
        if len(authHeader) > len("Bearer ") {
            tokenStr = authHeader[len("Bearer "):]
        }

        // If Authorization header is missing, check the token cookie
        if tokenStr == "" {
            cookie, err := r.Cookie("token")
            if err == nil {
                tokenStr = cookie.Value
            }
        }

        if tokenStr == "" {
            utils.Log.Warn("Missing authorization header or cookie")
            http.Error(w, "Missing authorization header or cookie", http.StatusUnauthorized)
            return
        }

        if utils.IsTokenBlacklisted(tokenStr) {
            utils.Log.Warn("Attempt to use blacklisted token")
            http.Error(w, "Token is blacklisted", http.StatusUnauthorized)
            return
        }

        claims := &models.Claims{}
        token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            return utils.JwtKey, nil
        })

        if err != nil || !token.Valid {
            utils.Log.WithFields(logrus.Fields{
                "event": "auth_middleware",
                "error": err,
            }).Error("Invalid token")
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Store the username in the request context (if needed)
        r.Header.Set("X-Authenticated-User", claims.Username)
        utils.Log.WithFields(logrus.Fields{
            "event":    "auth_middleware",
            "username": claims.Username,
        }).Info("User authenticated")

        // Call the next handler
        next.ServeHTTP(w, r)
    })
}

// RBAC middleware for role-based access control
func RBACMiddleware(requiredRole string, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenStr := r.Header.Get("Authorization")

        // Extract the user's role from the JWT
        userRole, err := utils.ExtractUserRoleFromToken(tokenStr)
        if err != nil || userRole != requiredRole {
            utils.Log.WithFields(logrus.Fields{
                "event":       "rbac_middleware",
                "requiredRole": requiredRole,
                "userRole":    userRole,
            }).Warn("Access denied due to insufficient privileges")
            http.Error(w, "Forbidden: insufficient privileges", http.StatusForbidden)
            return
        }

        utils.Log.WithFields(logrus.Fields{
            "event": "rbac_middleware",
            "role":  userRole,
        }).Info("Role-based access granted")


        // Call the next handler if the user has the required role
        next.ServeHTTP(w, r)
    })
}

// SecurityHeadersMiddleware adds security headers to the HTTP response
func SecurityHeadersMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Set security headers
        w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; object-src 'none';")
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-XSS-Protection", "1; mode=block")
        w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains") // For HTTPS

        // Call the next handler
        next.ServeHTTP(w, r)
    })
}

// ErrorHandlerMiddleware to log errors and prevent exposing stack traces
func ErrorHandlerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                utils.Log.WithFields(logrus.Fields{
                    "event": "panic",
                    "error": err,
                }).Error("A panic occurred during request processing")
                http.Error(w, "Internal server error", http.StatusInternalServerError)
            }
        }()
        next.ServeHTTP(w, r)
    })
}

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // Wrap the response writer to capture status code
        srw := &ResponseWriter{ResponseWriter: w, StatusCode: http.StatusOK}

        // Call the next handler
        next.ServeHTTP(srw, r)

        // Calculate response duration
        duration := time.Since(start).Seconds()

        // Record metrics
        utils.RequestCount.Add(1)
        utils.ResponseTime.Add(duration)
  

        // Log request details
        utils.Log.WithFields(logrus.Fields{
            "method":   r.Method,
            "path":     r.URL.Path,
            "status":   srw.StatusCode,
            "duration": time.Since(start),
        }).Info("Handled request")

        // Track error rates
        if srw.StatusCode >= 400 {
            utils.ErrorCount.Add(1) // Increment error count if status is 400 or higher
        }
    })
}

// responseWriter is a custom ResponseWriter to capture the status code
type ResponseWriter struct {
    http.ResponseWriter
    StatusCode int
}

func (rw *ResponseWriter) WriteHeader(code int) {
    rw.StatusCode = code
    rw.ResponseWriter.WriteHeader(code)
}


func CSRFProtection(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Skip CSRF validation for JSON requests
        if r.Header.Get("Content-Type") != "application/json" {
            // Apply CSRF protection
            csrf.Protect([]byte("32-byte-long-secret-key"))(next).ServeHTTP(w, r)
        } else {
            // Skip CSRF for JSON requests
            next.ServeHTTP(w, r)
        }
    })
}