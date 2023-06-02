package main

import (
	"log"
	"strconv"
	"sync"
	"time"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang/groupcache/lru"
)

var client *redis.Client
var mutex = &sync.Mutex{} 
var lruCache *lru.Cache
var hllKey = "ping_callers"


func init() {
	// Initialize Redis client
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
	})

	// Initialize LRU cache with a maximum of 10 entries
	lruCache = lru.New(10)
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
	username := c.PostForm("username")
	sessionID := generateSessionID()

	// Store the session ID and user name in Redis
	// err := client.HSet(client.Context(), "sessions", sessionID, username).Err()
	err := client.Set(c, sessionID, username, 0).Err()
	fmt.Println(err)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// c.JSON(200, gin.H{"username": username, "session_id": sessionID})
	c.JSON(http.StatusOK, gin.H{"session_id": sessionID})
}

func pingHandler(c *gin.Context) {
	sessionID := c.Query("session_id")

	// Check if the session ID is valid
	userName, err := client.Get(c, sessionID).Result()
	if err == redis.Nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid session ID"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve session"})
		return
	}

	// Acquire lock to ensure only one person can call /ping at a time
	mutex.Lock()
	defer mutex.Unlock()

	// Check if the user has exceeded the rate limit
	callCountKey := fmt.Sprintf("call_count:%s", userName)
	callCount, _ := lruCache.Get(callCountKey)
	if callCount == nil {
		callCount = 1
	} else {
		count, _ := strconv.Atoi(callCount.(string))
		if count >= 2 {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			return
		}
		callCount = count + 1
	}

	// Increment the call count
	lruCache.Add(callCountKey, callCount)

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
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("session:%d", timestamp)
}