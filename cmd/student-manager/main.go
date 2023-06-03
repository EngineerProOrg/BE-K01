package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/EngineerProOrg/BE-K01/configs"
	cache_service "github.com/EngineerProOrg/BE-K01/pkg/service/cache-service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"os"
)

var (
	confPath = flag.String("conf", "files/live.json", "path to config file")
)

func main() {
	config := &configs.StudentManagerConfig{}
	jsonFile, err := os.Open("configs/files/live.json")
	// if we os.Open returns an error then handle it
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
	err = json.Unmarshal(bt, config)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.DB.Addr,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Printf("can not connect to db :%v\n", err)
		return
	}

	rd := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: "",
		DB:       0,
	})
	if rd == nil {
		fmt.Printf("can not connect to redis :%v\n", err)
		return
	}

	r := gin.Default()

	cache_service.MappingService(r, cache_service.NewService(rd))

	r.Run("localhost:3050")
}
