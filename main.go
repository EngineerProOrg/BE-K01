package main

import(
	"github.com/gin-gonic/gin"
	"github.com/EngineerProOrg/BE-K01/internal/user/handler"
	"github.com/EngineerProOrg/BE-K01/tools/redis"
)
func main() {
	router := gin.Default()

	redis := redis.NewRedisClient()
	userHdl := handler.NewUserHandler(redis)
	router.POST("/login", userHdl.Login())
	router.GET("/ping", userHdl.Ping())
	router.GET("/get-counter", userHdl.GetCounter())
	router.GET("/rate-limit-ping", userHdl.RateLimitPing())
	router.GET("/top10", userHdl.Top10())
	router.GET("/count-hyperloglog", userHdl.CountHyperLogLog())
	router.Run(":8080")
}