package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Get PostgreSQL connection details from environment variables
	dbHost := os.Getenv("POSTGRES_HOST") // Change this if your database is hosted elsewhere
	dbPort := os.Getenv("POSTGRES_PORT") // Default PostgreSQL port
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	sslMode := os.Getenv("SSLMODE") // Adjust this based on your PostgreSQL setup

	// Construct the DATABASE_URL
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, sslMode)

	// Initialize the database connection
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal("connection error:", err)
		panic("Failed to connect to the database")
	}

	// Perform database operations here
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Image{})
	db.AutoMigrate(&Comment{})

	fmt.Println("Connected to PostgreSQL database")

	DB = db

}
