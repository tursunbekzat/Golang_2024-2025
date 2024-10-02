package main

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

type User struct {
    ID   uint   `gorm:"primaryKey"`
    Name string `gorm:"not null"`
    Age  int    `gorm:"not null"`
}

func connectDB() *gorm.DB {
    dsn := "host=localhost user=bekzat password=zx dbname=golang port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to the database:", err)
    }
    return db
}

func autoMigrate(db *gorm.DB) {
    err := db.AutoMigrate(&User{})
    if err != nil {
        log.Fatal("failed to auto-migrate:", err)
    }
}

func insertUser(db *gorm.DB, name string, age int) {
    user := User{Name: name, Age: age}
    result := db.Create(&user)

    if result.Error != nil {
        log.Fatal("Error inserting user:", result.Error)
    }
    log.Printf("Inserted user: %+v\n", user)
}

func queryUsers(db *gorm.DB) {
    var users []User
    result := db.Find(&users)

    if result.Error != nil {
        log.Fatal("Error querying users:", result.Error)
    }

    for _, user := range users {
        log.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
    }
}

func updateUser(db *gorm.DB, id uint, newName string) {
    var user User
    result := db.First(&user, id)

    if result.Error != nil {
        log.Fatal("User not found:", result.Error)
    }

    user.Name = newName
    db.Save(&user)
    log.Printf("Updated user: %+v\n", user)
}

func deleteUser(db *gorm.DB, id uint) {
    result := db.Delete(&User{}, id)

    if result.Error != nil {
        log.Fatal("Error deleting user:", result.Error)
    }
    log.Printf("Deleted user with ID: %d\n", id)
}

func main() {
    db := connectDB()
    defer func() {
        sqlDB, err := db.DB()
        if err != nil {
            log.Fatal(err)
        }
        sqlDB.Close()
    }()

    autoMigrate(db)

    // Вставка пользователей
    insertUser(db, "Alice", 25)
    insertUser(db, "Bob", 30)

    // Запрос пользователей
    queryUsers(db)

    // Обновление пользователя
    updateUser(db, 1, "Alice Smith")

    // Удаление пользователя
    deleteUser(db, 2)

    // Запрос пользователей после обновления и удаления
    queryUsers(db)
}
