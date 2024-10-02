package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/lib/pq"
)

// Connection string to connect to your PostgreSQL database
const (
    host     = "localhost"
    port     = 5432
    user     = "bekzat"
    password = "zx"
    dbname   = "golang"
)

func createTable(db *sql.DB) {
    query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT,
        age INT
    );`

    _, err := db.Exec(query)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Table 'users' created or already exists.")
}

func insertUser(db *sql.DB, name string, age int) {
    query := `
    INSERT INTO users (name, age) 
    VALUES ($1, $2)
    RETURNING id;`

    var id int
    err := db.QueryRow(query, name, age).Scan(&id)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Inserted user with ID: %d\n", id)
}

func queryUsers(db *sql.DB) {
    rows, err := db.Query("SELECT id, name, age FROM users")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    fmt.Println("Users in the database:")

    for rows.Next() {
        var id int
        var name string
        var age int

        err := rows.Scan(&id, &name, &age)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
    }

    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
}



func connectDB() *sql.DB {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }

    // Check if the connection works
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Successfully connected to PostgreSQL!")
    return db
}

func main() {
	db := connectDB()
    defer db.Close()

    // Create the 'users' table
    createTable(db)

    // Insert some users
    insertUser(db, "Alice", 25)
    insertUser(db, "Bob", 30)
    insertUser(db, "Charlie", 35)

    // Query and print all users
    queryUsers(db)
}
