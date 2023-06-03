package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var redisClient *redis.Client
var router *gin.Engine
var db *gorm.DB

var key_ping string

type User struct {
	ID       int    `json:"-" gorm:"column:id;"`
	Username string `json:"username" gorm:"column:name;"`
	Password string `json:"password" gorm:"column:password;"`
}

func (User) TableName() string {
	return "User"
}

func init() {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:12346@tcp(127.0.0.1:3307)/todo_list?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println("can not connect to db ", err)
		return
	}

	// Khởi tạo Redis Client
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	router = gin.Default()

	key_ping = "ping_ping"
}

func get_data(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data User
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = db.Where("id =? ", id).First(&data).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})

	}

}

func login() {
	// Đăng ký API /login
	router.POST("/login", func(c *gin.Context) {

		// Lấy thông tin đăng nhập từ request body
		var loginInfo struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		var data User
		if err := c.BindJSON(&loginInfo); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		// query from db
		err := db.Where(&User{Username: loginInfo.Username, Password: loginInfo.Password}).First(&data).Error

		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid username or password"})
			return
		}

		// Tạo session id ngẫu nhiên
		rand.Seed(time.Now().UnixNano())
		sessionID := strconv.Itoa(rand.Intn(1000000000))

		// Lưu thông tin session vào Redis
		value := fmt.Sprintf("%d+%s+%s", data.ID, data.Username, data.Password)
		err_redis := redisClient.Set(redisClient.Context(), sessionID, value, time.Hour).Err()
		if err_redis != nil {
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}

		// Trả về thông tin đăng nhập thành công và session id
		c.JSON(200, gin.H{"message": "Login successful", "session_id": sessionID})
	})
}

var flag bool

func ping() {
	// Đăng ký API /ping

	router.GET("/ping", func(c *gin.Context) {
		if flag {
			c.AbortWithStatusJSON(429, gin.H{"error": "too many requests"})
			return
		}
		sessionID := c.GetHeader("sessionID")

		value, err := redisClient.Get(redisClient.Context(), sessionID).Result()
		if err != nil {
			c.AbortWithStatusJSON(429, gin.H{"error": "cant get data user"})
		}
		parts := strings.Split(value, "+")
		fmt.Println(parts)
		// Tăng biến đếm lên 1 và lưu trữ vào Redis

		mapKey := fmt.Sprintf("user:%s:MapCount ", parts[0])

		mapValue, errPing := redisClient.HGetAll(redisClient.Context(), mapKey).Result()
		now := time.Now().Unix()
		timeCount, err := strconv.ParseInt(mapValue["timeCount"], 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "convert time false"})
			return
		}

		if errPing != nil || now-timeCount > 60 {
			myMap := map[string]interface{}{
				"timeCount": now,
				"count":     0,
			}
			errMap := redisClient.HSet(redisClient.Context(), mapKey, myMap).Err()
			if errMap != nil {
				c.AbortWithStatusJSON(500, gin.H{"error": "set key map error"})
				return
			}
		}

		numPing, errNumPing := strconv.ParseInt(mapValue["timeCount"], 10, 64)

		if errNumPing != nil || numPing >= 2 {
			c.AbortWithStatusJSON(500, gin.H{"error": "user already ping 2 time"})
			return
		}
		// increase count
		errIncreaseCount := redisClient.HIncrBy(redisClient.Context(), mapKey, "count", 1).Err()
		if errIncreaseCount != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "increase count in 60s"})
			return
		}

		// đếm số lần ping của user
		keyUser := fmt.Sprintf("user:%s:request_count", parts[0])
		_, err1 := redisClient.Incr(redisClient.Context(), keyUser).Result()
		if err1 != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "internal server error"})
			return
		}

		// Truy vấn số lượng request của người dùng đó từ Redis
		count, err := redisClient.Get(redisClient.Context(), keyUser).Int()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "get total count error"})
			return
		}

		count60, errCount60s := redisClient.HGet(redisClient.Context(), mapKey, "count").Result()

		if errCount60s != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "get count60 error"})
			return
		}

		//
		numPings, err := redisClient.ZScore(redisClient.Context(), key_ping, keyUser).Result()

		c.JSON(200, gin.H{"all count": count, "count in 60s": count60, "all count with key root": numPings})

		flag = true
		// Thực hiện sleep trong 5 giây

		time.Sleep(5 * time.Second)
		flag = false
	})
}

func topPing() {
	router.GET("/top", func(c *gin.Context) {
		// Lấy top 10 người gọi API /ping nhiều nhất
		result, err := redisClient.ZRevRangeWithScores(redisClient.Context(), key_ping, 0, 9).Result()
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Internal server error",
			})
			return
		}

		// Định dạng kết quả trảvề cho người dùng
		var topUsers []gin.H
		for _, z := range result {
			topUsers = append(topUsers, gin.H{
				"username":   z.Member,
				"ping_calls": z.Score,
			})
		}

		// Trả về kết quả
		c.JSON(200, topUsers)
	})
}

func hyperLogLog() {

	router.GET("/count", func(c *gin.Context) {
		// Get the unique API call count from Redis
		count, err := redisClient.PFCount(redisClient.Context(), key_ping).Result()
		if err != nil {
			// Handle error
		}

		// Return the count to the user
		c.JSON(200, gin.H{
			"count": count,
		})
	})
}
func main() {

	login()
	ping()
	topPing()
	hyperLogLog()
	router.Run(":8080")
}
