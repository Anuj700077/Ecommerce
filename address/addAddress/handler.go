package address

import (
	"Ecommerce/address"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddAddressHandler(c *gin.Context) {

	var addr address.Address

	if err := c.ShouldBindJSON(&addr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// get user_id from middleware
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// ✅ FIXED LINE
	addr.UserID = uint(userID.(int))

	err := AddAddressService(&addr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address added successfully"})
}
