package handler

import (
	"github.com/EngineerProOrg/BE-K01/internal/user/model"
	"github.com/EngineerProOrg/BE-K01/tools/redis"
	"github.com/gin-gonic/gin"
)

// đếm số lượng lần 1 người gọi api /ping
func (hdl *userHandler) GetCounter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userName model.LoginPing
		if err := ctx.ShouldBindJSON(&userName); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		redisClient := redis.NewRedisClient()
		count := redisClient.Get(ctx, userName.Username).Val()
		ctx.JSON(200, gin.H{"count": count})
	}
}