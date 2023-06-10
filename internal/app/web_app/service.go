package web_app

import (
	v1 "github.com/EngineerProOrg/BE-K01/internal/app/web_app/v1"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	v1Router := r.Group("v1")
	v1.AddUserRouter(v1Router)
	r.Run(":8080")
}
