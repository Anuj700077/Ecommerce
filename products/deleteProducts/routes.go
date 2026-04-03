package Products

import (
	"Ecommerce/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(rg *gin.RouterGroup) {

	admin := rg.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())

	{
		admin.DELETE("/product/:id", DeleteProductHandler)
	}
}
