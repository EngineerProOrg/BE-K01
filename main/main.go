package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/gorilla/securecookie"
	"log"
	"net/http"
	"strings"
	"time"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func main() {
	_, err := redisClient.Ping(redisClient.Context()).Result()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	// 1. API /login, để tạo session cho mỗi người đăng nhập, dùng redis để lưu session id, user name ấy
	router.POST("/login", loginHandler)
	// 2. API /ping chỉ cho phép 1 người được gọi tại một thời điểm ( với sleep ở bên trong api đó trong 5s)
	router.GET("/ping", pingHandler)
	// 3. API đếm số lượng lần 1 người gọi api /ping
	router.GET("/ping-count/:usernam", countNumberOfRequestHandler)
	// 4. rate limit mỗi người chỉ được gọi API /ping 2 lần trong 60s
	router.GET("/rate-limiter/:username", rateLimiterHandler)

	router.Run(":8080")
}

func rateLimiterHandler(c *gin.Context) {
	username := c.Params.ByName("username")
	redisClient.Incr(redisClient.Context(), username)
	requestNumberOfUser, _ := redisClient.Get(redisClient.Context(), username).Int64()
	if requestNumberOfUser == 1 {
		redisClient.Expire(redisClient.Context(), username, 60*time.Second)
	}

	if requestNumberOfUser >= 3 {
		fmt.Println("exceeded the number of requests")
		c.JSON(http.StatusInternalServerError, "exceeded the number of requests")
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func countNumberOfRequestHandler(c *gin.Context) {
	username := c.Params.ByName("username")
	redisClient.Incr(redisClient.Context(), username)
	numberOfRequest := redisClient.Get(redisClient.Context(), username).Val()
	jsonResult := "number of request is :" + numberOfRequest
	c.JSON(http.StatusOK, jsonResult)
}

func pingHandler(c *gin.Context) {
	newUUID := uuid.New().String()
	log.Println(" thread : ", newUUID)
	key := "redisLock"
	expireTime := 10 * time.Minute // TTL of key
	sleepTime := 10 * time.Second
	tryLockRedisForKey(key, newUUID, expireTime)
	sleep(sleepTime)
	// doSomething

	if strings.Compare(redisClient.Get(c, key).Val(), newUUID) == 0 {
		releaseRedisLockForKey(key)
	}

	c.JSON(http.StatusOK, "ping successfully")

}

func sleep(sleepTime time.Duration) {
	time.Sleep(sleepTime)
}

func tryLockRedisForKey(key string, value string, expireTime time.Duration) {
	for {
		result := redisClient.SetNX(redisClient.Context(), key, value, expireTime)
		if result.Err() != nil {
			log.Fatal(" setNX failed")
			return
		}
		if result.Val() {
			break
		}

	}
}

func releaseRedisLockForKey(key string) {
	err := redisClient.Del(redisClient.Context(), key).Err()
	if err != nil {
		log.Fatal("del key failed ")
		return
	}
}
func isRedisLockForKey(key string, value string) bool {
	return redisClient.SetNX(redisClient.Context(), key, value, 10*time.Minute).Val()
}

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(c *gin.Context) {
	var dataItem LoginInfo

	if err := c.BindJSON(&dataItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	username := dataItem.Username
	sessionId := generateSession()
	err := redisClient.Set(c, sessionId, username, 1*time.Minute).Err()
	fmt.Println(err)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, "login successfully")

}

func generateSession() string {
	return string(securecookie.GenerateRandomKey(32))
}
