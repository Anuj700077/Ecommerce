package Users

import "github.com/gin-gonic/gin"

func UserRoutes(r *gin.RouterGroup) {

	if r == nil {
		return
	}

	r.POST("/send-otp", SendOTP)
	r.POST("/verify-otp", VerifyOTPAndRegisterHandler)
}
