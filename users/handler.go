package Users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendOTP(c *gin.Context) {

	var req OTPRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format. Please send proper JSON."})
		return
	}

	err := SendOTPService(req.Email)
	if err != nil {

		switch err {

		case ErrEmailRequired:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required. Please provide your email address."})

		case ErrInvalidEmail:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format. Please enter a valid email."})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong. Please try again later."})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

// Verify OTP + Register Handler
func VerifyOTPAndRegisterHandler(c *gin.Context) {

	var req VerifyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format. Please send proper JSON."})
		return
	}

	user := User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
	}

	err := VerifyOTPAndRegister(&user, req.OTP)
	if err != nil {

		switch err {

		case ErrFieldsMissing:
			c.JSON(http.StatusBadRequest, gin.H{"error": " Please fill all details."})

		case ErrInvalidEmail:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format."})

		case ErrInvalidOTP:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OTP. Please check...."})

		case ErrOTPExpired:
			c.JSON(http.StatusBadRequest, gin.H{"error": "OTP has expired. Please request a new one."})

		case ErrOTPNotFound:
			c.JSON(http.StatusBadRequest, gin.H{"error": "No OTP found. Please request OTP first."})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed. Please try again later."})
		}

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func LoginHandler(c *gin.Context) {

	var req LoginRequest

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

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
	})
}
