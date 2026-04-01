package userlogin

import "github.com/gin-gonic/gin"

func Userlogin(r *gin.RouterGroup) {

	r.POST("/login", LoginHandler)

}
