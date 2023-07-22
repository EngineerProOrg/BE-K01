package v1

import (
	"github.com/EngineerProOrg/BE-K01/internal/app/web_app/service"
	"github.com/gin-gonic/gin"
)

func AddNewsfeedRouter(r *gin.RouterGroup, svc *service.WebService) {
	routerGroup := r.Group("newsfeeds")

	routerGroup.GET("", svc.GetNewsfeed)
}
