package address

import (
	"net/http"
	"strconv"

	"Ecommerce/address"

	"github.com/gin-gonic/gin"
)

func UpdateAddressHandler(c *gin.Context) {

	var addr address.Address


	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address id"})
		return
	}


	if err := c.ShouldBindJSON(&addr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}


	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	addr.ID = uint(id)
	addr.UserID = uint(userID.(int))

	err = UpdateAddressService(&addr)

	if err != nil {

		if err.Error() == "address not found or unauthorized" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address updated successfully"})
}
