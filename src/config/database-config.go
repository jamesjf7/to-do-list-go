package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // connect to database

	if err != nil {
		// The panic() function is inbuilt in Go Language and when it is raised, the code prints a panic message, and the function crashes.
		panic("failed to connect database")
	}
	fmt.Println("Successfully connected to database!")
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("failed to close database")
	}
	dbSQL.Close()
}
