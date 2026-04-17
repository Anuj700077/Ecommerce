package address

import (
	"Ecommerce/middleware"

	"github.com/gin-gonic/gin"
)

func AddressRoutes(r *gin.RouterGroup) {

	auth := r.Group("/address")
	auth.Use(middleware.AuthMiddleware())

	auth.GET("/", GetAddressHandler)
}
