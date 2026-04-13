package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB   // GORM
var SQLDB *sql.DB // RAW SQL

func ConnectDatabase() {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	//  Step 1: Connect to default postgres DB
	psqlInfo := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=postgres port=%s sslmode=%s",
		host, user, password, port, sslmode,
	)

	sqlDB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Failed to connect PostgreSQL:", err)
	}

	//  Check connection
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Database not reachable:", err)
	}

	//  Create DB if not exists
	_, err = sqlDB.Exec("CREATE DATABASE " + dbname)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		log.Fatal("Failed to create database:", err)
	}

	sqlDB.Close()

	//  Step 2: Connect to actual DB using GORM
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect DB:", err)
	}

	DB = db

	// 🔹 Get RAW SQL DB
	sqlDB2, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get SQL DB:", err)
	}

	SQLDB = sqlDB2

	//  Connection pool (important)
	SQLDB.SetMaxOpenConns(10)
	SQLDB.SetMaxIdleConns(5)

	log.Println("Database connected successfully")
}

func CreateTables() {

	if SQLDB == nil {
		log.Fatal("SQLDB not initialized")
	}

	// USERS TABLE
	userQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		phone TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		role TEXT DEFAULT 'user'
	);
	`

	_, err := SQLDB.Exec(userQuery)
	if err != nil {
		log.Fatal("Failed to create users table:", err)
	}

	// PRODUCTS TABLE
	productQuery := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		price INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = SQLDB.Exec(productQuery)
	if err != nil {
		log.Fatal("Failed to create products table:", err)
	}

	// ADDRESS TABLE (Linked with Users)
	addressQuery := `
	CREATE TABLE IF NOT EXISTS addresses (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		address1 TEXT NOT NULL,
		pincode TEXT NOT NULL,
		city TEXT NOT NULL,
		country TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		CONSTRAINT fk_user
			FOREIGN KEY(user_id)
			REFERENCES users(id)
			ON DELETE CASCADE
	);
	`

	_, err = SQLDB.Exec(addressQuery)
	if err != nil {
		log.Fatal("Failed to create addresses table:", err)
	}

	log.Println("Tables created successfully")
}
