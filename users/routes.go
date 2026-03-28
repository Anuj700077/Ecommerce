package users

import "github.com/gin-gonic/gin"

func UserRoutes(r *gin.RouterGroup) {

	if r == nil {
		return
	}

	r.POST("/register", Register)
}
