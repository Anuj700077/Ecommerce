package Products

import (
	"Ecommerce/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(rg *gin.RouterGroup) {

	//here user can see the products without login 
	products := rg.Group("/products")
	{
		products.GET("", GetAllProducts)
	}

//this route is for admin
	admin := rg.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())

	{
		admin.POST("/product", CreateProducts) 
	}
}
