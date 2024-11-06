package models

import  (
    "github.com/dgrijalva/jwt-go"
)

// User represents the user structure
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username" validate:"required"`
    Password string `json:"password" validate:"required"`
    Role     string `json:"role" validate:"required"`
}

// Claims struct to store JWT claims
type Claims struct {
    Username string `json:"username"`
    Role     string `json:"role"`
    jwt.StandardClaims
}