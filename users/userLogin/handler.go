package userlogin

import (
	"Ecommerce/Users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {

	var req Users.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
		})
		return
	}

	user, err := LoginService(req.Email, req.Password)

	if err != nil {

		switch err {

		case ErrFieldsMissing:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password required"})

		case ErrInvalidEmail:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})

		case ErrUserNotFound:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Firstly go and signup"})

		case ErrInvalidPassword:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Login failed"})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}
