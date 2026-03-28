package main

import (
	"log"

	database "Ecommerce/database"
	"Ecommerce/router"

	"github.com/joho/godotenv"
)

func main() {

	// Load ENV
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	log.Println("Environment variables loaded")

	// Connect Database
	database.ConnectDatabase()

	database.CreateTables()
	// Setup Router
	r := router.SetupRouter()

	// Start Server
	router.StartServer(r)
}
