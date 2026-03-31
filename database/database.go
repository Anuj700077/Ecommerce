package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB      // GORM (optional)
var SQLDB *sql.DB   // ✅ RAW SQL (IMPORTANT)

func ConnectDatabase() {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	// Step 1: Connect to default postgres DB
	psqlInfo := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=postgres port=%s sslmode=%s",
		host, user, password, port, sslmode,
	)

	sqlDB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Failed to connect PostgreSQL:", err)
	}

	// Create DB if not exists
	_, err = sqlDB.Exec("CREATE DATABASE " + dbname)
	if err != nil {
		log.Println("Database may already exist:", err)
	}

	sqlDB.Close()

	// Step 2: Connect to your DB using GORM
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect DB:", err)
	}

	DB = db

	//  IMPORTANT: Get raw SQL DB
	sqlDB2, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get SQL DB:", err)
	}

	SQLDB = sqlDB2

	log.Println("Database connected successfully")
}

func CreateTables() {

	if SQLDB == nil {
		log.Fatal("SQLDB not initialized")
	}

	// Users table
	userQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		phone TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);
	`

	_, err := SQLDB.Exec(userQuery)
	if err != nil {
		log.Fatal("Failed to create users table:", err)
	}

	//  Products table
	productQuery := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		price INT NOT NULL
	);
	`

	_, err = SQLDB.Exec(productQuery)
	if err != nil {
		log.Fatal("Failed to create products table:", err)
	}

	log.Println("Tables created successfully")
}
