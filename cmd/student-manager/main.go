package main

import (
	"fmt"
	"net/http"

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

type Cource struct {
	CourseId   int32  `gorm:"course_id"`
	CourseName string `gorm:"course_name"`
}

func responseError(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err,
	})
}

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
	rd := redis.NewClient(&redis.Options{
		Addr:     "localhost:6377",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err != nil {
		fmt.Println("can not connect to db ", err)
		return
	}

	if rd == nil {
		fmt.Println("connect fail to redis")
		return
	}

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		if db != nil {
			fmt.Println("connected DB ")
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/professor", getProfessors(db))
	r.GET("/courses", getCourceDistinct(db))

	r.Run()
}
