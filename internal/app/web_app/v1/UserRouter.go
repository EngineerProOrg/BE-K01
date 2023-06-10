package v1

import (
	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	userRouter := r.Group("users")
	userRouter.GET("", func(context *gin.Context) {

	})
	userRouter.POST("", func(context *gin.Context) {

	})
	userRouter.PUT("", func(context *gin.Context) {

	})
}
