package adminlogin

import "github.com/gin-gonic/gin"

func UserLogin(r *gin.RouterGroup) {
	r.POST("/admin/login", AdminLoginHandler)
}
