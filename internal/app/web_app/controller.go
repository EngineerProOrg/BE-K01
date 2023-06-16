package web_app

import (
	"fmt"

	"github.com/EngineerProOrg/BE-K01/internal/app/web_app/service"
	v1 "github.com/EngineerProOrg/BE-K01/internal/app/web_app/v1"
	"github.com/gin-gonic/gin"
)

type WebController struct {
	service.WebService
	Port int
}

func (c WebController) Run() {
	r := gin.Default()
	v1Router := r.Group("v1")
	v1.AddUserRouter(v1Router, &c.WebService)

	r.Run(fmt.Sprintf(":%d", c.Port))
}
