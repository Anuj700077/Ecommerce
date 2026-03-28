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

var DB *gorm.DB

func ConnectDatabase() {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	
	psqlInfo := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=postgres port=%s sslmode=%s",
		host, user, password, port, sslmode,
	)

	sqlDB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Failed to connect to PostgreSQL server: %v", err)
	}


	_, err = sqlDB.Exec("CREATE DATABASE " + dbname)
	if err != nil {
		log.Println("Database may already exist or error occurred:", err)
	} else {
		log.Println("Database created successfully")
	}

	sqlDB.Close()

	
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Unable to connect to database: %v", err)
	}

	DB = db

	log.Println("Connected to ecommerce database ")
}


func CreateTables() {

	if DB == nil {
		log.Fatal("Database not connected. Cannot create tables")
	}

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		phone TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);
	`

	err := DB.Exec(query).Error
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

	log.Println("Users table is ready")
}
