package Products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddProducts(c *gin.Context){

}




func GetProducts(c *gin.Context) {

	_, err := GetProductsService()
	if err != nil {

		if err == ErrNoProductsFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "No products available",})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products",})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Products fetched succesfully",
	})
}
