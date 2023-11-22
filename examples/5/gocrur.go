package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    int
	Name  string
	Email string
}

func main() {
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect user database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create id
	db.Create(&User{ID: 1, Name: "xxx", Email: "yyy@zzz"})

	// Read id
	var user User
	db.First(&user, 1)

	// Update it
	db.Model(&user).Update("name", "yyy")

	// Read it gain
	db.First(&user, 1)

	// Print it
	fmt.Printf("User: %+v\n", user)
}
