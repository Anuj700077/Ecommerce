package Products

import "github.com/gin-gonic/gin"

func ProductRoutes(rg *gin.RouterGroup) {

	products := rg.Group("/products")

	
	products.GET("", GetProducts)
	products.POST("",AddProducts)
}
