package Products

import (
	"Ecommerce/Products"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateProductHandler(c *gin.Context) {

	var product Products.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := UpdateProductService(&product)
	if err != nil {

		switch err {

		case ErrInvalidProduct:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})

		case ErrProductNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully",
	})
}
