package router

import (
	adminlogin "Ecommerce/Users/adminLogin"
	otpandregister "Ecommerce/Users/otpAndRegister"
	userlogin "Ecommerce/Users/userLogin"

	createProducts "Ecommerce/Products/createProducts"
	editProducts "Ecommerce/Products/editProducts"
	getProducts "Ecommerce/Products/getProducts"

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
		c.JSON(200, gin.H{"message": "Server is running "})
	})

	api := r.Group("/api")

	otpandregister.UserRoutes(api)
	userlogin.Userlogin(api)
	adminlogin.UserLogin(api)

	createProducts.ProductRoutes(api)
	getProducts.ProductRoutes(api)
	editProducts.ProductRoutes(api)

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
