package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=karzal dbname=booksdb port=3000 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	log.Println("Database connected successfully.")
}
