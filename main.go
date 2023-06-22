package main

import (
	"log"
	"time"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var client *redis.Client
var hllKey = "ping_callers"


func init() {
	// Initialize Redis client
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
	})
}

func main() {

	// Check to connect Redis
	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.POST("/login", loginHandler)
	router.GET("/ping", pingHandler)
	router.GET("/top", topHandler)
	router.GET("/count", countHandler)

	router.Run(":8080")
}

func loginHandler(c *gin.Context) {
	sessionID := c.GetHeader("session_id")
	if sessionID != "" {
		// Session ID already exists, no need to create a new session
		c.JSON(http.StatusOK, gin.H{"message": "Already logged in"})
		return
	}

	username := c.PostForm("username")
	sessionID = generateSessionID()

	// Store the session ID and user name in Redis
	err := client.Set(c, sessionID, username, 0).Err()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Set sessionID cookie
	c.SetCookie("session_id", sessionID, 300, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"session_id": sessionID, "username": username, "message": "Session created successfully"})
}

func pingHandler(c *gin.Context) {
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "session ID not found"})
		return
	}

	// Check if the session ID is valid
	userName, err := client.Get(c, sessionID).Result()
	if err == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid session ID"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve session"})
		return
	}

	// Acquire distributed lock to ensure only one person can call /ping at a time
	lockKey := fmt.Sprintf("lock:%s", userName)
	ok, err := client.SetNX(c, lockKey, "locked", 5*time.Second).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to obtain lock"})
		return
	}
	if !ok {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "failed to obtain lock"})
		return
	}

	// Implement Redis-based rate limiting
	rateLimitKey := fmt.Sprintf("ratelimit:%s", userName)
	remaining, err := client.Incr(c, rateLimitKey).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to increment rate limit"})
		return
	}

	if remaining > 2 {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
		return
	}

	// Sleep for 5 seconds to simulate processing time
	time.Sleep(5 * time.Second)

	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func topHandler(c *gin.Context) {
	// Retrieve the top 10 most-called /ping APIs
	result, err := client.ZRevRangeWithScores(c, hllKey, 0, 9).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve top APIs"})
		return
	}

	var topAPIs []string
	for _, z := range result {
		topAPIs = append(topAPIs, z.Member.(string))
	}

	c.JSON(http.StatusOK, gin.H{"top_apis": topAPIs})
}

func countHandler(c *gin.Context) {
	// Retrieve the approximate number of callers to /ping API using HyperLogLog
	count, err := client.PFCount(c, hllKey).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve count"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}

func generateSessionID() string {
	return uuid.New().String()
}