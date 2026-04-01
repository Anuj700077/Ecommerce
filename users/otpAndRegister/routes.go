package otpandregister

import "github.com/gin-gonic/gin"

func UserRoutes(r *gin.RouterGroup) {

	r.POST("/send-otp", SendOTP)
	r.POST("/verify-otp", VerifyOTPAndRegisterHandler)
	
	

}
