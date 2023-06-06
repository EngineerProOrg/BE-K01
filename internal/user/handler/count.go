package handler

import (
	"github.com/EngineerProOrg/BE-K01/tools/redis"
	"github.com/gin-gonic/gin"
)

//Dùng hyperloglog để lưu xấp sỉ số người gọi api /ping , và trả về trong api /count

func (hdl *userHandler) CountHyperLogLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		redisClient := redis.NewRedisClient()
		count := redisClient.PFCount(ctx, "membersCallPingAPI").Val()
		ctx.JSON(200, gin.H{"count": count})
	}
	
}