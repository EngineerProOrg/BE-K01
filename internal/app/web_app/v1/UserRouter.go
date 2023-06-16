package v1

import (
	"github.com/EngineerProOrg/BE-K01/internal/app/web_app/service"
	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup, svc *service.WebService) {
	userRouter := r.Group("users")
	userRouter.GET("", svc.CheckUserNamePassword)
	userRouter.POST("", svc.CreateUser)
	userRouter.PUT("", svc.EditUser)
}
