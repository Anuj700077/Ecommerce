package address

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAddressHandler(c *gin.Context) {

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not Authorized"})
		return
	}
	uid := uint(userID.(int))

	addresses, err := GetAddressService(uid)
	if err != nil {
		if err.Error() == "no address found" {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "FAiled to fetch Address"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"addresses": addresses,
	})
}
