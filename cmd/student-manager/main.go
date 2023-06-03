package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Professor struct {
	ProfId    int32  `gorm:"prof_id"`
	ProfLname string `gorm:"prof_lname"`
	ProfFname string `gorm:"prof_fname"`
	StudLname string `gorm:"stud_lname"`
	StudFname string `gorm:"stud_fname"`
	NumClass  int32  `gorm:"num_class"`
}

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Cource struct {
	CourseId   int32  `gorm:"course_id"`
	CourseName string `gorm:"course_name"`
}

func responseError(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err,
	})
}

var rd = redis.NewClient(&redis.Options{
	Addr:     "localhost:6377",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func getProfessors(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data []Professor

		err := db.Table("professors as p").
			Select("p.prof_id, p.prof_lname, p.prof_fname, s.stud_lname, s.stud_fname,COUNT(*) as num_class").
			Joins("JOIN classes as c ON c.prof_id = p.prof_id").
			Joins("JOIN enrolls as e ON c.class_id = e.class_id").
			Joins("JOIN students as s ON s.stud_id = e.stud_id").
			Group("p.prof_id, s.stud_id").
			Find(&data).Error
		if err != nil {
			responseError(c, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

func getCourceDistinct(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data []Cource
		err := db.Table("courses as c").
			Select("c.*").
			Joins("JOIN classes as cl ON c.course_id = cl.course_id").
			Joins("JOIN professors as p ON p.prof_id = cl.prof_id").
			Where("p.prof_id = ? ", 1).
			Find(&data).Error

		if err != nil {
			responseError(c, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if len(username) == 0 || len(password) == 0 {
		// handle validate
		fmt.Println("Please enter username and password")
	}

	if username != "admin" || password != "1234" {
		fmt.Println("Incorrect !!")
		return
	}
	key := "user_" + username
	ctx := context.Background()
	err := rd.SetNX(ctx, key, username, time.Duration(0)).Err()

	if err != nil {
		fmt.Println("Fail !!")
		return
	}
	fmt.Println("Save session done !!")
}

var keyApiPing = "user_accessing"

// API /ping chỉ cho phép 1 người được gọi tại một thời điểm ( với sleep ở bên trong api đó trong 5s)
func ping(c *gin.Context) {
	ctx := context.Background()

	exists, err := rd.Exists(ctx, keyApiPing).Result()
	if err != nil {
		fmt.Println("error ", err)
		return
	}

	// check nếu đang tồn tại key thì sleep  5 - ttl của key
	if exists == 1 {
		ttl, err := rd.TTL(ctx, keyApiPing).Result()
		if err != nil {
			fmt.Println("Lỗi khi lấy TTL:", err)
			return
		}
		fmt.Printf("ttl %d", ttl)
		if ttl > 0 {
			waitTime := time.Duration(ttl)
			fmt.Printf("wait %d minutes", ttl)
			time.Sleep(waitTime)
		}
	} else {
		// khi vào, tạo một key trong redis set time 5s
		err := rd.SetNX(ctx, keyApiPing, "1", 5*time.Second).Err()
		if err != nil {
			return
		}
	}

	// return resource
	fmt.Println("Access resource")
}

func main() {
	// connect to mysql
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:12345@tcp(127.0.0.1:3305)/engineerpro?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	// connect to redis

	if err != nil {
		fmt.Println("can not connect to db ", err)
		return
	}

	if rd == nil {
		fmt.Println("connect fail to redis")
		return
	}

	r := gin.Default()

	r.GET("/professor", getProfessors(db))
	r.GET("/courses", getCourceDistinct(db))

	// Bai tap redis
	r.POST("/login", login)
	r.GET("/ping", ping)

	r.Run()
}
