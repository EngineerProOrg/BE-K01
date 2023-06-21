package web_app

import (
	"fmt"
	"net/http/pprof"

	"github.com/EngineerProOrg/BE-K01/internal/app/web_app/service"
	v1 "github.com/EngineerProOrg/BE-K01/internal/app/web_app/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type WebController struct {
	service.WebService
	Port int
}

func (c WebController) Run() {
	r := gin.Default()

	v1Router := r.Group("/api/v1")
	v1.AddUserRouter(v1Router, &c.WebService)
	initSwagger(r)
	initPprof(r)
	r.Run(fmt.Sprintf(":%d", c.Port))
}

func initSwagger(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func initPprof(r *gin.Engine) {
	r.GET("/debug/pprof/", func(context *gin.Context) {
		pprof.Index(context.Writer, context.Request)
	})
	r.GET("/debug/pprof/profile", func(context *gin.Context) {
		pprof.Profile(context.Writer, context.Request)
	})
	r.GET("/debug/pprof/trace", func(context *gin.Context) {
		pprof.Trace(context.Writer, context.Request)
	})
}
