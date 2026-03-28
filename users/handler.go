package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	// here calling service (logics)
	err := RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to register user. Please check input."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user": user})
}
