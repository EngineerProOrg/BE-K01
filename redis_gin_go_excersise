package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang/groupcache/lru"
)

var (
	redisClient *redis.Client
)

// khởi tạo 1 biến mutex (khóa)
// Mutex là cơ chế đồng bộ hóa đảm bảo chỉ 1 gorountine (luồng)
// được thực hiện trong 1 thời điểm
var mutex = &sync.Mutex{}

var lruCache *lru.Cache

// 1 func để khởi tạo ra 1 redis Client mới
// và thông báo error nếu có
// func không có trả về
// mà thay đổi trực tiếp global variable: redisClient
func initRedis() {
	// return a new client to the redis server
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	// sử dụng Ping để xem server có chạy và kết nối thành công hay không
	// .Result() để trả lại kết quả của lệnh Ping
	// đưa context.Context vào trong lệnh đễ provide context
	_, err := redisClient.Ping(redisClient.Context()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis", err)
	}

	// cho lru cache voi max la 10 entries
	lruCache = lru.New(10)
}
func generateSessionID() string {
	// tạo ra 1 session ID ngẫu nhiên
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
func logInHandler(c *gin.Context) {
	// lấy ra key từ trong POST urlencoded form
	username := c.PostForm("username")
	// password := c.PostForm("password")
	// gin.H là 1 cái map <=> key-value, vs key là string
	// JSON: serialize struct thành JSON trong response body
	// ShouldBindJSON sẽ bind struct pointer và xem nó có thành JSON đc không
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Invalid request", // không được thì coi như request k hợp lệ
	// 	})
	// 	return
	// }
	sessionID := generateSessionID()
	// HSet được sử dụng để lưu trữ một cặp khóa-giá trị trong một hash (bảng băm) của Redis
	// khóa là sessionID còn "username"->field, username->value
	// err := redisClient.HSet(redisClient.Context(), sessionID, "username", username, "password", password).Err()
	err := redisClient.Set(redisClient.Context(), sessionID, username, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session"})
		return
	}
	// username, _ = redisClient.HGet(redisClient.Context(), sessionID, "username").Result()
	// password, _ = redisClient.HGet(redisClient.Context(), sessionID, "password").Result()
	// lưu
	c.JSON(http.StatusOK, gin.H{
		"message":    "Login successful",
		"session_id": sessionID,
		"username":   username,
	})
}

func main() {
	initRedis()
	// tạo ra 1 Object Router để định nghĩa
	// các route xử lý HTTP
	// có sẵn middleware:
	// logger: ghi lại http và thông tin như phg thức đường dẫn
	// recovery: xử lý panic, lỗi, đbảo sever không crash
	router := gin.Default()
	router.POST("/login", logInHandler)
	router.GET("/ping", pingHandler)
	router.GET("/top", topUserHandler)
	router.GET("/count", countHandler)
	// run attaches the router to a http.Server
	// start listening and serving http requests

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

}
func pingHandler(c *gin.Context) {
	// the sessionID will be saved as session_id
	sessionID := c.Query("session_id")
	// log.Printf("SessionID: %T\n", sessionID)
	// Check if the session ID is valid
	// look up in the redis cached
	// Result() sẽ trả về giá trị value với key cung cấp và err
	userName, err := redisClient.Get(redisClient.Context(), sessionID).Result()
	// nếu không thấy có trong redis cache thì => sessionID k hợp lệ
	if err == redis.Nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":    "invalid session ID",
			"username": userName,
		})
		return
	} else if err != nil {
		// nếu err có trả về gì đó
		// có lỗi với hệ thống
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve session",
		})
		return
	}

	// đặt khóa để chỉ 1 người /ping được 1 lúc mà thôi
	mutex.Lock()
	defer mutex.Unlock() // mở lock sau khi thực hiện hết

	// tạo ra callCountKey cho lru.Cache lấy để get trong cache
	callCountKey := fmt.Sprintf("call_count:%s", userName)
	// Chúng ta sử dụng lru.Cache để lưu trữ và lấy ra số lần gọi của mỗi user
	// sử dụng redis là nơi lưu chính cho session info(id, username)
	// nhưng dùng lru.Cache để lưu call count giúp cải thiện perf
	// The LRU cache acts as a local cache that holds a subset of the data in memory and provides faster access.
	// giảm số redis query
	// lru.Cache.Get() sẽ lấy khóa callCountKey có dạng call_count:<userName
	// rồi trả về value của key đó nếu không trả về nil
	callCount, _ := lruCache.Get(callCountKey)
	// nếu không có value của key đó
	if callCount == nil {
		// số lần gọi không có trong cache
		// => ch được gọi bao giờ
		// đây là lần 1
		callCount = 1
	} else {
		// đổi từ string sang int
		count, _ := strconv.Atoi(callCount.(string))
		if count >= 2 {
			// nếu đây là ping lần thứ >= 2 => không cho
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			return
		} else {
			// nếu không thì thêm 1 vào coi như thêm 1 lần ping
			callCount = count + 1
		}
	}
	// add lần gọi ping này vào hyperloglog
	err = redisClient.PFAdd(redisClient.Context(), "myset", sessionID).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to increment ping count",
		})
		return
	}
	// sau khi đã cập nhật callCount
	// chúng ta sẽ lưu nó lại vào cache tiếp
	// với key là: callCountKey và val: callCount
	lruCache.Add(callCountKey, callCount)

	// Sleep với 5 giây để mô phỏng thời gian xử lý
	time.Sleep(5 * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"message": "ping successfully",
	})
}

func topUserHandler(c *gin.Context) {
	// sử dụng hàm ZRevRangeWithScores
	//trả về một danh sách các thành viên trong một Sorted Set
	// theo thứ tự giảm dần (theo score) cùng với các điểm số tương ứng
	sortedSet, err := redisClient.ZRevRangeWithScores(redisClient.Context(), "myset", 0, 9).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve top 10 users calling API",
		})
	}
	var topUsers []string
	for _, val := range sortedSet {
		// lấy phần tử trong sortedSet = val.Member
		// rồi ép kiểu về string cho vào slice
		topUsers = append(topUsers, val.Member.(string))
	}
	c.JSON(http.StatusOK, gin.H{
		"top_users": topUsers,
	})
}

func countHandler(c *gin.Context) {
	// hàm này sử dụng để trả về
	// số ng gọi API
	// sử dụng hyperloglog cho kết quả gần đúng
	count, err := redisClient.PFCount(redisClient.Context(), "myset").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve count"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}
