package cache_service

import (
	"github.com/EngineerProOrg/BE-K01/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"time"
)

type CacheService interface {
	Login(userName string) (string, error)
	Ping(sessionId string) error
	Top10Ping() ([]string, error)
	Count() (int64, error)
	CountBySessionId(sessionId string) (int64, error)
}

type cacheService struct {
	redis *redis.Client
}

func NewService(redis *redis.Client) CacheService {
	return &cacheService{
		redis: redis,
	}
}

func MappingService(r *gin.Engine, service CacheService) {
	r.POST("/login", func(c *gin.Context) {
		var payload types.LoginPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		sessionId, err := service.Login(payload.UserName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"sessionId": sessionId})
	})

	r.POST("/ping", func(c *gin.Context) {
		var payload types.RequestPayload
		err := c.ShouldBindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = service.Ping(payload.SessionID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		time.Sleep(time.Second * 5)

		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	r.GET("/top", func(c *gin.Context) {
		result, err := service.Top10Ping()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"scoreboard": result})
	})

	r.GET("/count", func(c *gin.Context) {
		result, err := service.Count()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"count": result})
	})

	r.GET("/count/:sessionId", func(c *gin.Context) {
		sessionId := c.Param("sessionId")
		result, err := service.CountBySessionId(sessionId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"count": result})
	})
}
