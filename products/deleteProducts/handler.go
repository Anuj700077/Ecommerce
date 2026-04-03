package Products

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteProductHandler(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	err = DeleteProducts(id)
	if err != nil {
		switch err {
		case ErrInvalidID:
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Product ID"})
		case ErrProductNotFound:
			c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete product"})
		}
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": "Product Deleted Successfully"})
}
