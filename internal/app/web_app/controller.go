package web_app

import (
	"fmt"
	"net/http/pprof"

	"github.com/EngineerProOrg/BE-K01/internal/app/web_app/service"
	v1 "github.com/EngineerProOrg/BE-K01/internal/app/web_app/v1"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	v1.AddFriendRouter(v1Router, &c.WebService)
	initSwagger(r)
	initPprof(r)
	initPrometheus(r)
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
func initPrometheus(r *gin.Engine) {
	handler := promhttp.Handler()
	r.GET("/metrics", func(context *gin.Context) {
		handler.ServeHTTP(context.Writer, context.Request)
	})
}
