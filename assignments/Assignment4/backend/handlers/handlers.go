package handlers

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
    "github.com/gorilla/csrf"
    "github.com/sirupsen/logrus"
)

// Handler for the home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the Home Page!")
}

// Register handler for user registration and form display
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    if r.Method == http.MethodPost {
        var user models.User

        // Check if the request is JSON
        if r.Header.Get("Content-Type") == "application/json" {
            // Parse JSON body
            if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
                utils.Log.WithFields(logrus.Fields{
                    "event": "registration_attempt",
                    "error": err.Error(),
                }).Error("Failed to parse registration data")
                http.Error(w, "Invalid JSON data", http.StatusBadRequest)
                return
            }
        } else {
            // Parse form data
            r.ParseForm()
            user.Username = r.FormValue("username") // Ensure correct field names
            user.Password = r.FormValue("password") // Ensure correct field names
            user.Role = r.FormValue("role")         // Optional: If you want to set user role during registration
        }

        // Validate the input
        if err := utils.Validate.Struct(user); err != nil {
            utils.Log.WithFields(logrus.Fields{
                "event":   "registration_attempt",
                "username": user.Username,
                "error":   err.Error(),
            }).Warn("Validation failed for registration")
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        // Hash the password securely
        hashedPassword, err := utils.HashPassword(user.Password)
        if err != nil {
            utils.Log.WithFields(logrus.Fields{
                "event":   "registration_attempt",
                "username": user.Username,
                "error":   err.Error(),
            }).Error("Password hashing failed")
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            return
        }

        // Store the user
        utils.Users[user.Username] = hashedPassword
        
        // Log successful registration
        utils.Log.WithFields(logrus.Fields{
            "event":    "registration_success",
            "username": user.Username,
        }).Info("User registered successfully")

        // Redirect to login page after successful registration
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    // Render the registration form (for HTML requests)
    tmpl := `
    <form method="POST" action="/register">
        <input type="hidden" name="_csrf" value="{{.}}">
        Username: <input type="text" name="username" required><br>
        Password: <input type="password" name="password" required><br>
        Role: <input type="text" name="role" required placeholder="user/admin"><br>
        <input type="submit" value="Register">
    </form>`
    csrfToken := csrf.Token(r)
    t, _ := template.New("form").Parse(tmpl)
    t.Execute(w, csrfToken)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    if r.Method == http.MethodPost {
        var user models.User
        err := r.ParseForm()
        if err != nil {
            utils.LogError("Failed to parse form data", err)
            http.Error(w, "Invalid request", http.StatusBadRequest)
            return
        }
        // Check if the request is JSON
        if r.Header.Get("Content-Type") == "application/json" {
            // Parse JSON body
            err := json.NewDecoder(r.Body).Decode(&user)
            if err != nil {
                utils.Log.WithFields(logrus.Fields{
                    "event": "login_attempt",
                    "error": err.Error(),
                }).Error("Failed to parse login request")
                http.Error(w, "Invalid JSON data", http.StatusBadRequest)
                return
            }
        } else {
            // Parse form data
            r.ParseForm()
            user.Username = r.FormValue("username") // Ensure correct field names
            user.Password = r.FormValue("password") // Ensure correct field names
        }

        // Check if the user exists
        storedPassword, exists := utils.Users[user.Username]
        if !exists || utils.VerifyPassword(storedPassword, user.Password) != nil {
            utils.Log.WithFields(logrus.Fields{
                "event":    "login_attempt",
                "username": user.Username,
            }).Warn("Failed login attempt with invalid credentials")
            http.Error(w, "Invalid username or password", http.StatusUnauthorized)
            return
        }

        // Create the JWT claims, which includes the username
        expirationTime := time.Now().Add(15 * time.Minute)
        claims := &models.Claims{
            Username: user.Username,
            Role:     user.Role,
            StandardClaims: jwt.StandardClaims{
                ExpiresAt: expirationTime.Unix(),
            },
        }

        // Create the JWT token
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
        tokenString, err := token.SignedString(utils.JwtKey)
        if err != nil {
            utils.Log.WithFields(logrus.Fields{
                "event": "login_attempt",
                "error": err.Error(),
            }).Error("Failed to create JWT token for user")
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            return
        }

        utils.Log.WithFields(logrus.Fields{
            "event":    "login_attempt",
            "username": user.Username,
        }).Info("User logged in successfully")

        // Set the token in a cookie if not JSON request
        if r.Header.Get("Content-Type") != "application/json" {
            http.SetCookie(w, &http.Cookie{
                Name:     "token",
                Value:    tokenString,
                Expires:  expirationTime,
                HttpOnly: true,
                Secure:   false, // Set to true in production with HTTPS
            })
            fmt.Fprintf(w, "Login successful! JWT token set in cookie.")
        } else {
                // For JSON responses, send the token in the response body
                w.Header().Set("Content-Type", "application/json")
                json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
        }   
    
        // Redirect to a protected resource
        http.Redirect(w, r, "/protected", http.StatusSeeOther)
        return
    }

    // Render the login form (for HTML requests)
    tmpl := `
    <form method="POST" action="/login">
        <input type="hidden" name="_csrf" value="{{.}}">
        Username: <input type="text" name="username" required><br>
        Password: <input type="password" name="password" required><br>
        <input type="submit" value="Login">
    </form>`
    csrfToken := csrf.Token(r)
    t, _ := template.New("form").Parse(tmpl)
    t.Execute(w, csrfToken)
}

// Logout handler to clear user session
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    tokenStr := r.Header.Get("Authorization")

    // Check if token is provided and has the correct prefix
    if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Remove "Bearer " prefix
    tokenStr = tokenStr[len("Bearer "):]

    // Retrieve username from token (assuming a utility function to do this)
    username, err := utils.GetUsernameFromToken(tokenStr)
    if err != nil {
        utils.Log.WithFields(logrus.Fields{
            "event": "logout_attempt",
            "error": err.Error(),
        }).Warn("Failed to retrieve username from token during logout")
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }

    // Add token to blacklist
    utils.AddTokenToBlacklist(tokenStr)

    // Log successful logout
    utils.Log.WithFields(logrus.Fields{
        "event":    "logout",
        "username": username,
    }).Info("User logged out successfully")

    // Redirect to login page after logging out
    http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// ProtectedHandler for a protected route
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("This is a protected area"))
}
