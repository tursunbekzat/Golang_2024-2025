package utils

import (
	"database/sql"
	_ "github.com/lib/pq"
	// "fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "log"
)

func ConnectSQL() (*sql.DB, error) {
	connStr := "postgres://user:password@localhost/dbname?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ConnectGORM() (*gorm.DB, error) {
	dsn := "user=bekzat password=zx dbname=golang port=5432 sslmode=disable TimeZone=Asia/Almaty"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
