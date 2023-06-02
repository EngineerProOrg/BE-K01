package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/EngineerProOrg/BE-K01/configs"
	// "github.com/EngineerProOrg/BE-K01/pkg/controller"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	confPath = flag.String("conf", "configs/files/live.json", "path to config file")
)

func main() {
	// Read configuration file from `confPath`
	jsonFile, err := os.Open(*confPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()
	bt, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Parse contents of configuration file to `config`
	config := &configs.StudentManagerConfig{}
	err = json.Unmarshal(bt, config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize database
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "quangmx:2511@tcp(127.0.0.1:3306)/engineerpro?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println("Can not connect to db:", err)
		return
	}

	// Initialize redis
	redisClient := redis.NewClient(&redis.Options{})
	if redisClient == nil {
		fmt.Println("Can not initialize redis")
		return
	}

	// // service := service.NewService(db, rd)
	// // controller.MappingService(r, service)
	// s := service.NewService(db, redisClient)
	// service.MappingService(router, s)
	// router.Run()

	// Initialize gin
	router := gin.Default()

	// Declare /login API
	router.POST("/login", handleLogin)

	// Declare /ping API
	router.GET("/ping", handlePing)

	// Delcare /top API
	router.GET("/top", handleTop)

	// Declare /count API
	router.GET("/count", handleCount)

	// Start the web server
	err = router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
