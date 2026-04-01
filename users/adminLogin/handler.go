package adminlogin

import (
	"Ecommerce/Users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminLoginHandler(c *gin.Context) {

	var req Users.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	err := AdminLoginService(req.Email, req.Password)

	if err != nil {

		switch err {

		case ErrFieldsMissing:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password required"})

		case ErrAdminInvalidCredentials:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid admin credentials"})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Admin login failed"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Admin login successful"})
}