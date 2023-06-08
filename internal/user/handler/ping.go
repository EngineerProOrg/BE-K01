package handler

import (
	"github.com/EngineerProOrg/BE-K01/internal/user/model"
	"github.com/gin-gonic/gin"
	"time"
)

// 1 API /ping chỉ cho phép 1 người được gọi tại một thời điểm ( với sleep ở bên trong api đó trong 5s)
func (hdl *userHandler) Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var nameUser model.LoginPing
		if err := ctx.ShouldBindJSON(&nameUser); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		
		lockKey := "distributed-lock"
		lockValue := "distributed-lock-value"

		setNXResult := hdl.redis.SetNX(ctx, lockKey, lockValue, time.Duration(5000000000))
		if setNXResult.Err() != nil || !setNXResult.Val() {
			ctx.JSON(400, gin.H{"error": "lock is being used"})
			return
		}

		
		hdl.redis.Incr(ctx, nameUser.Username) // serve as a counter for the number of requests of each user

		hdl.redis.ZIncrBy(ctx, "leaderboard", 1, nameUser.Username) // serve for api top10 (leaderboard)

		hdl.redis.PFAdd(ctx, "membersCallPingAPI", nameUser.Username) // serve for api count unique users call api /ping use hyperloglog

		ctx.JSON(200, gin.H{"message": "pong"})
	}
}
