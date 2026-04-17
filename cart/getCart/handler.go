package cart

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCartHandler(c *gin.Context) {


	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	cartItems, total, err := GetCartService(uint(userID.(int)))

	if err != nil {

		if err.Error() == "cart is empty" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cart_items": cartItems,
		"grand_total": total,
	})
}
