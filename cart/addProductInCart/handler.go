package cart

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func AddToCartHandler(c *gin.Context) {

	var req Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err := AddToCartService(uint(userID.(int)), req.ProductID, req.Quantity)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})
}
