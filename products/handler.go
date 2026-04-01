package Products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProducts(c *gin.Context) {

	var product Product

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

func GetAllProducts(c *gin.Context) {

	products, err := GetProduct()
	if err != nil {

		if err == ErrNoProductsFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "No products available"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}
