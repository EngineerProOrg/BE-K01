package handler

import (
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
}

func NewUserHandler() UserHandler {
	return &userHandler{}
}