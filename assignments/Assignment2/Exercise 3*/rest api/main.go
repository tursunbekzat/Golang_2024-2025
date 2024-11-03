package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"rest-api/handlers"
	"rest-api/utils"
)

func main() {
	sqlDB, err := utils.ConnectSQL()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	gormDB, err := utils.ConnectGORM()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	// Routes for SQL handlers
	router.HandleFunc("/sql/users", handlers.GetUsersSQL(sqlDB)).Methods("GET")
	router.HandleFunc("/sql/users", handlers.CreateUserSQL(sqlDB)).Methods("POST")

	// Routes for GORM handlers
	router.HandleFunc("/gorm/users", handlers.GetUsersGORM(gormDB)).Methods("GET")
	router.HandleFunc("/gorm/users", handlers.CreateUserGORM(gormDB)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
