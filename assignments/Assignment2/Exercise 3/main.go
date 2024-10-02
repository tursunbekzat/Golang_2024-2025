package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
)

var db *sql.DB

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func connectDB() {
    var err error
    dsn := "host=localhost user=bekzat password=zx dbname=golang port=5432 sslmode=disable"
    db, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, name, age FROM users")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        users = append(users, user)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    query := "INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id"
    err := db.QueryRow(query, user.Name, user.Age).Scan(&user.ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    query := "UPDATE users SET name = $1, age = $2 WHERE id = $3"
    _, err := db.Exec(query, user.Name, user.Age, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    query := "DELETE FROM users WHERE id = $1"
    _, err := db.Exec(query, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func main() {
    connectDB()
    defer db.Close()

    r := mux.NewRouter()
    r.HandleFunc("/users", getUsers).Methods("GET")
    r.HandleFunc("/user", createUser).Methods("POST")
    r.HandleFunc("/user/{id}", updateUser).Methods("PUT")
    r.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8080", r))
}
