package models

type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"unique;not null" json:"name"`
	Age   int    `gorm:"not null" json:"age"`
	Email string `gorm:"unique;not null" json:"email"`
}
