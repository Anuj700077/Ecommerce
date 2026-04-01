package Products

import (
	"Ecommerce/Products"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProducts(c *gin.Context) {

	var product Products.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := CreateProductService(&product)
	if err != nil {

		switch err {

		case ErrInvalidProduct:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product name and valid price required"})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
	})
}
