package handler

import (
	"github.com/redis/go-redis/v9"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Login() gin.HandlerFunc
	Ping() gin.HandlerFunc
	GetCounter() gin.HandlerFunc
	RateLimitPing() gin.HandlerFunc
	Top10() gin.HandlerFunc
	CountHyperLogLog() gin.HandlerFunc

}
type userHandler struct {
	redis *redis.Client
}

func NewUserHandler(redis *redis.Client) UserHandler {
	return &userHandler{redis: redis}
}