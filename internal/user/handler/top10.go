package handler

import (
	"github.com/EngineerProOrg/BE-K01/tools/redis"
	"github.com/gin-gonic/gin"
)

func (hdl *userHandler) Top10() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		redisClient := redis.NewRedisClient()
		top10, err := redisClient.ZRevRangeWithScores(ctx, "leaderboard", 0, 9).Result()
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"top10": top10})
	}
}