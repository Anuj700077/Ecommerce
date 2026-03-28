package router

import (
	users "Ecommerce/Users"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is running ",
		})
	})

	api := r.Group("/api")

	users.UserRoutes(api)
	return r
}

func StartServer(r *gin.Engine) {

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("PORT not set in .env, defaulting to 8080")
		port = "8080"
	}

	log.Printf("Server starting on port %s...\n", port)

	if err := r.Run(":" + port); err != nil {
		log.Print("Failed to start server: %v", err)
	}
}
