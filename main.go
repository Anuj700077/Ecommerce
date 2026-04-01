package main

import (
	"log"

	database "Ecommerce/database"
	"Ecommerce/router"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	log.Println("Environment variables loaded")

	database.ConnectDatabase()
	database.CreateTables()

	r := router.SetupRouter()
	router.StartServer(r)
}
