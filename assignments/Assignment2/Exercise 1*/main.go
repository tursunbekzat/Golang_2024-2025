package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

// User struct for fetching from the database
type User struct {
	ID   int // ID field for the user
	Name string
	Age  int
}

// UserInput struct for inserting into the database (without ID)
type UserInput struct {
	Name string
	Age  int
}

func main() {
	var err error
	connStr := "user=bekzat dbname=golang password=zx sslmode=disable" // Adjust this
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	// Check if the connection is alive
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable()

	// Use UserInput for inserting
	users := []UserInput{
		{"Bauka", 30},
		{"Maks", 25},
		{"Bauka", 30}, // Trying to insert duplicate
	}

	insertUsers(users)

	// Query users without filter
	queryUsers(nil, 1, 10)

	// Update a user
	updateUser(1, "Alice Updated", 31)

	// Delete a user
	deleteUser(2)
}

// Create the users table
func createTable() {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) UNIQUE NOT NULL,
			age INT NOT NULL
		);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// Insert multiple users into the users table
func insertUsers(users []UserInput) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		err := insertUser(tx, user)
		if err != nil {
			fmt.Printf("Error inserting user %s: %v\n", user.Name, err)
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

// Insert a single user with duplicate check
func insertUser(tx *sql.Tx, user UserInput) error {
	var id int
	err := tx.QueryRow("SELECT id FROM users WHERE name = $1", user.Name).Scan(&id)
	if err == sql.ErrNoRows {
		// User does not exist, proceed with insertion
		_, err := tx.Exec("INSERT INTO users (name, age) VALUES ($1, $2)", user.Name, user.Age)
		if err != nil {
			return err
		}
		fmt.Printf("Inserted user: %s\n", user.Name)
	} else if err != nil {
		return err
	} else {
		fmt.Printf("User %s already exists with ID %d\n", user.Name, id)
	}
	return nil
}

func queryUsers(ageFilter *int, page, pageSize int) {
    query := "SELECT id, name, age FROM users"
    var args []interface{}

    // Add age filter if provided
    if ageFilter != nil {
        query += " WHERE age = $1"
        args = append(args, *ageFilter)
    }

    query += " LIMIT $1 OFFSET $2"
    // Если возрастной фильтр используется, параметры будут $2 и $3
    // Если возрастной фильтр не используется, параметры будут $1 и $2
    if ageFilter != nil {
        args = append(args, pageSize, (page-1)*pageSize) // Если фильтр есть
    } else {
        args = append(args, pageSize, (page-1)*pageSize) // Если фильтра нетp
    }

    // Prepare the statement
    stmt, err := db.Prepare(query)
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    // Execute the statement with arguments
    rows, err := stmt.Query(args...)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Name, &user.Age) // Scanning ID
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("User: %+v\n", user)
    }

    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
}


// Update a user’s details
func updateUser(id int, name string, age int) {
	_, err := db.Exec("UPDATE users SET name = $1, age = $2 WHERE id = $3", name, age, id)
	if err != nil {
		log.Fatal(err)
	}
}

// Delete a user by their ID
func deleteUser(id int) {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}
