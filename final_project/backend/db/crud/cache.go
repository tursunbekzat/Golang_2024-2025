package crud

import (
	"fmt"
	"backend/db"
	"time"
)

// Create Cache
func CreateCache(key, value string, expiration time.Time) {
	cache := db.Cache{CacheKey: key, CacheValue: value, ExpirationTime: expiration}
	result := db.DB.Create(&cache)
	if result.Error != nil {
		fmt.Println("Error creating cache:", result.Error)
		return
	}
	fmt.Println("Cache created successfully:", cache)
}

// Get Cache by Key
func GetCache(key string) {
	var cache db.Cache
	result := db.DB.First(&cache, "cache_key = ?", key)
	if result.Error != nil {
		fmt.Println("Cache not found:", result.Error)
		return
	}
	fmt.Println("Cache:", cache)
}

// Delete Cache
func DeleteCache(key string) {
	result := db.DB.Delete(&db.Cache{}, "cache_key = ?", key)
	if result.Error != nil {
		fmt.Println("Error deleting cache:", result.Error)
		return
	}
	fmt.Println("Cache deleted successfully.")
}
