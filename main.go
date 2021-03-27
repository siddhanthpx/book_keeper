package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Book struct {
	Name       string
	Author     string
	CallNumber int `gorm:"unique_index"`
	PersonID   int
}

type Person struct {
	gorm.Model
	Name  string
	Email string `gorm:"typevarchar100;unique_index"`
	Books []Book
}

func main() {
	err := godotenv.Load()
	checkError(err)

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	// Database Connection
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, name, port)
	// Opening Connction to DB
	db, err := gorm.Open(postgres.Open(dbURI))
	checkError(err)

	log.Println("Successfully connected to database.", db)
}
