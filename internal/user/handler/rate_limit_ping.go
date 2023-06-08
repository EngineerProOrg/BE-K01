package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/EngineerProOrg/BE-K01/internal/user/model"
	"time"
)

// rate limit mỗi người chỉ được gọi API /ping 2 lần trong 60s

func (hdl *userHandler) RateLimitPing() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var nameUser model.LoginPing
		if err := ctx.ShouldBindJSON(&nameUser); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		exists, _ := hdl.redis.Exists(ctx, nameUser.Username).Result()
		if exists == 0 {
			hdl.redis.Set(ctx, nameUser.Username, 0, time.Duration(60*time.Second))
		} else {
			count := hdl.redis.Get(ctx, nameUser.Username).Val()
			if count == "2" {
				ctx.JSON(400, gin.H{"error": "limit 2 requests per 5s"})
				return
			}
		}
		hdl.redis.Incr(ctx, nameUser.Username) // serve as a counter for the number of requests of each user

		ctx.JSON(200, gin.H{"message": "pong"})
	}
}