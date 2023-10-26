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

func ConnectDB(secretConf configs.SecretsConfig) {

	dbHost := secretConf.PostgresHost
	dbPort := secretConf.PostgresPort
	dbName := secretConf.PostgresDb
	dbUser := secretConf.PostgresUser
	dbPassword := secretConf.PostgresPassword
	sslMode := secretConf.Sslmode

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
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&Account{})
	// db.AutoMigrate(&Post{})
	// db.AutoMigrate(&Rent{})
	// db.AutoMigrate(&Image{})
	// db.AutoMigrate(&Booking{})
	// db.AutoMigrate(&Location{})
	// db.AutoMigrate(&Review{})
	// db.AutoMigrate(&Transaction{})
	// db.AutoMigrate(&Cat{}) // for testing purpose !!!
	// db.AutoMigrate(&ForgotPassword{})

	// Perform database operations here with error handling
	if err := autoMigrateModel(db, &User{}); err != nil {
		log.Fatalf("Error migrating User: %v", err)
	}
	if err := autoMigrateModel(db, &Account{}); err != nil {
		log.Fatalf("Error migrating Account: %v", err)
	}
	if err := autoMigrateModel(db, &Post{}); err != nil {
		log.Fatalf("Error migrating Post: %v", err)
	}
	if err := autoMigrateModel(db, &Rent{}); err != nil {
		log.Fatalf("Error migrating Rent: %v", err)
	}
	if err := autoMigrateModel(db, &Image{}); err != nil {
		log.Fatalf("Error migrating Image: %v", err)
	}
	if err := autoMigrateModel(db, &Booking{}); err != nil {
		log.Fatalf("Error migrating Booking: %v", err)
	}
	if err := autoMigrateModel(db, &Location{}); err != nil {
		log.Fatalf("Error migrating Location: %v", err)
	}
	if err := autoMigrateModel(db, &Review{}); err != nil {
		log.Fatalf("Error migrating Review: %v", err)
	}
	if err := autoMigrateModel(db, &Transaction{}); err != nil {
		log.Fatalf("Error migrating Transaction: %v", err)
	}
	if err := autoMigrateModel(db, &Cat{}); err != nil {
		log.Fatalf("Error migrating Cat: %v", err)
	}
	if err := autoMigrateModel(db, &ForgotPassword{}); err != nil {
		log.Fatalf("Error migrating ForgotPassword: %v", err)
	}

	fmt.Println("Connected to PostgreSQL database")

	DB = db

}

func autoMigrateModel(db *gorm.DB, model interface{}) error {
	if err := db.AutoMigrate(model); err != nil {
		return err
	}
	return nil
}
