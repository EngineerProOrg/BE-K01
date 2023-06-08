package handler

import (
	"github.com/gin-gonic/gin"
)

//Dùng hyperloglog để lưu xấp sỉ số người gọi api /ping , và trả về trong api /count

func (hdl *userHandler) CountHyperLogLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		count := hdl.redis.PFCount(ctx, "membersCallPingAPI").Val()
		ctx.JSON(200, gin.H{"count": count})
	}
	
}