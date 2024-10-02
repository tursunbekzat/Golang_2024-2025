package main

import (
	"fmt"
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var db *gorm.DB

// User model
type User struct {
    ID      uint    `gorm:"primaryKey"`
    Name    string  `gorm:"not null"`
    Age     int     `gorm:"not null"`
    Profile Profile
}

// Profile model
type Profile struct {
    ID               uint   `gorm:"primaryKey"`
    UserID           uint   `gorm:"uniqueIndex;constraint:OnDelete:CASCADE;"` // Automatically deletes Profile when User is deleted
    Bio              string
    ProfilePictureURL string
}

func initDB() {
    var err error
    dsn := "host=localhost user=bekzat password=zx dbname=golang port=5432 sslmode=disable"
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to database:", err)
    }

    sqlDB, err := db.DB()
    if err != nil {
        log.Fatal(err)
    }

    sqlDB.SetMaxOpenConns(10)
    sqlDB.SetMaxIdleConns(5)
}

func createUserWithProfile() {
    user := User{
        Name: "Bekzat",
        Age:  20,
        Profile: Profile{
            Bio:              "Software Developer",
            ProfilePictureURL: "http://example.com/profile.jpg",
        },
    }

    result := db.Create(&user)
    if result.Error != nil {
        log.Fatal("failed to create user and profile:", result.Error)
    }

    log.Println("User and Profile created successfully")
}

func getUsersWithProfiles() {
    var users []User

    result := db.Preload("Profile").Find(&users)
    if result.Error != nil {
        log.Fatal("failed to fetch users with profiles:", result.Error)
    }

    for _, user := range users {
        log.Printf("User: %s, Age: %d, Bio: %s, Profile Picture: %s\n",
            user.Name, user.Age, user.Profile.Bio, user.Profile.ProfilePictureURL)
    }
}

func updateUserProfile(userID uint, newBio string, newPicture string) {
    var profile Profile
    result := db.First(&profile, "user_id = ?", userID)
    if result.Error != nil {
        log.Fatal("failed to find profile:", result.Error)
    }

    profile.Bio = newBio
    profile.ProfilePictureURL = newPicture

    result = db.Save(&profile)
    if result.Error != nil {
        log.Fatal("failed to update profile:", result.Error)
    }

    log.Println("Profile updated successfully")
}

func deleteUser(db *gorm.DB, userID uint) error {
    // First, delete the associated profile
    if err := db.Where("user_id = ?", userID).Delete(&Profile{}).Error; err != nil {
        return fmt.Errorf("failed to delete profile: %v", err)
    }

    // Then, delete the user
    if err := db.Delete(&User{}, userID).Error; err != nil {
        return fmt.Errorf("failed to delete user: %v", err)
    }

    fmt.Println("User and associated profile deleted successfully")
    return nil
}


func main() {
    initDB()
    
    err := db.AutoMigrate(&User{}, &Profile{})
    if err != nil {
        log.Fatal("failed to migrate database:", err)
    }

    createUserWithProfile()
    getUsersWithProfiles()
    updateUserProfile(1, "Updated Bio", "http://example.com/updated.jpg")
    deleteUser(db, 1)
}
