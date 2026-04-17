package cart

import (
	"Ecommerce/middleware"

	"github.com/gin-gonic/gin"
)

func CartRoutes(r *gin.RouterGroup) {

	auth := r.Group("/cart")
	auth.Use(middleware.AuthMiddleware())

	auth.GET("/", GetCartHandler)
}
