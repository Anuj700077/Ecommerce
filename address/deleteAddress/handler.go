package address

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteAddressHandler(c *gin.Context) {

	// 🔹 get id from URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address id"})
		return
	}

	// 🔹 get user_id from middleware
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err = DeleteAddressService(uint(id), uint(userID.(int)))

	if err != nil {

		if err.Error() == "address not found or unauthorized" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete address"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address deleted successfully"})
}
