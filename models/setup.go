package models

import (
	"fmt"
	"log"

	// "github.com/joho/godotenv"
	"github.com/risqiikhsani/rentvehicles/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(secret configs.SecretsConfig) {

	// err := godotenv.Load()
	// if err != nil {
	// 	panic("Error loading .env file")
	// }

	// Get PostgreSQL connection details from environment variables
	dbHost := secret.PostgresHost // Change this if your database is hosted elsewhere
	dbPort := secret.PostgresPort // Default PostgreSQL port
	dbName := secret.PostgresDb
	dbUser := secret.PostgresUser
	dbPassword := secret.PostgresPassword
	sslMode := secret.Sslmode // Adjust this based on your PostgreSQL setup

	// Construct the DATABASE_URL
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, sslMode)

	// Initialize the database connection
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		// https://gorm.io/docs/constraints.html
		DisableForeignKeyConstraintWhenMigrating: true,
		// https://gorm.io/docs/performance.html
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		log.Fatal("connection error:", err)
		panic("Failed to connect to the database")
	}

	// Perform database operations here
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Rent{})
	db.AutoMigrate(&Image{})
	db.AutoMigrate(&Booking{})
	db.AutoMigrate(&Location{})
	db.AutoMigrate(&Review{})
	db.AutoMigrate(&Transaction{})
	db.AutoMigrate(&Cat{}) // for testing purpose !!!

	fmt.Println("Connected to PostgreSQL database")

	DB = db

}
