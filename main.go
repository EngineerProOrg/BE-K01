package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

}

func main() {
	_, err := redisClient.Ping(redisClient.Context()).Result()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.POST("/login", loginHandler)
	router.POST("/ping", pingHandler)
	router.POST("/ping-count", pingCountHandler)
	router.POST("/rate-limit", rateLimitHandler)
	router.POST("/top", topHandler)
	router.POST("/count", countHandler)
	router.POST("/hyperloglog", hyperloglogHandler)
	router.Run(":8080")

}

func hyperloglogHandler(c *gin.Context) {
	count, err := redisClient.PFCount(redisClient.Context(), "hyperloglog").Result()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}

func topHandler(c *gin.Context) {
	top, err := redisClient.ZRevRangeWithScores(redisClient.Context(), "top", 0, 9).Result()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, top)
}

func countHandler(context *gin.Context) {

}

func rateLimitHandler(c *gin.Context) {
	cookie, err := c.Request.Cookie("session_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//check cookie exist in redis
	username, err := redisClient.Get(redisClient.Context(), cookie.Value).Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	count, err := redisClient.Incr(redisClient.Context(), username).Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if count == 1 {
		redisClient.Expire(redisClient.Context(), username, time.Minute)
	}

	if count >= 3 {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func generateSessionID() string {
	sessionToken := uuid.NewString()
	return sessionToken
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func checkLogin(loginRequest LoginRequest) bool {
	return loginRequest.Username == "admin" && loginRequest.Password == "admin"
}

func loginHandler(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.BindJSON(&loginRequest); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !checkLogin(loginRequest) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	sessionToken := generateSessionID()
	fmt.Println(sessionToken)

	err := redisClient.Set(redisClient.Context(), sessionToken, loginRequest.Username, time.Minute).Err()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(time.Minute),
	})
	c.JSON(http.StatusOK, gin.H{"message": "login success"})
}

func acquireLock(redisClient *redis.Client, lockKey string, value string, lockTimeout time.Duration) bool {
	for {
		result := redisClient.SetNX(redisClient.Context(), lockKey, value, lockTimeout)
		if result.Err() != nil {
			log.Fatal("Failed to acquire lock")
			return false
		}
		if result.Val() {
			return true
		}

		// Lock not acquired, wait and retry
		time.Sleep(time.Second)
	}
}

func releaseLock(redisClient *redis.Client, lockKey string) bool {
	result := redisClient.Del(redisClient.Context(), lockKey)
	if result.Err() != nil {
		log.Fatal("Failed to release lock")
		return false
	}
	return true
}

func pingHandler(c *gin.Context) {
	cookie, err := c.Request.Cookie("session_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//check cookie exist in redis
	username, err := redisClient.Get(redisClient.Context(), cookie.Value).Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//thread
	value := uuid.NewString()
	fmt.Sprintf("thread: %s", value)

	lockKey := "lock"
	if !acquireLock(redisClient, lockKey, value, 5*time.Minute) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to acquire lock"})
		return
	}
	defer releaseLock(redisClient, lockKey)
	time.Sleep(5 * time.Second)
	redisClient.Incr(redisClient.Context(), username)                 //count
	redisClient.ZIncrBy(redisClient.Context(), "top", 1, username)    //top
	redisClient.PFAdd(redisClient.Context(), "hyperloglog", username) //hyperloglog
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello %s", username)})
}

func pingCountHandler(c *gin.Context) {
	cookie, err := c.Request.Cookie("session_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check cookie exist in redis
	username, err := redisClient.Get(redisClient.Context(), cookie.Value).Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//count
	count, err := redisClient.Get(redisClient.Context(), username).Int()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Count: %d", count)})
}
