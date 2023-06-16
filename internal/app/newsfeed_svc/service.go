package newsfeed_svc

import (
	"context"
	"fmt"

	"github.com/EngineerProOrg/BE-K01/configs"
	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/newsfeed"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type NewsfeedService struct {
	newsfeed.UnimplementedNewsfeedServer
	db    *gorm.DB
	redis *redis.Client
}

func (n NewsfeedService) Newsfeed(ctx context.Context, request *newsfeed.NewsfeedRequest) (*newsfeed.NewsfeedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewNewsfeedService(conf *configs.NewsfeedConfig) (*NewsfeedService, error) {
	db, err := gorm.Open(mysql.New(conf.MySQL), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println("can not connect to db ", err)
		return nil, err
	}
	rd := redis.NewClient(&conf.Redis)
	if rd == nil {
		return nil, fmt.Errorf("can not init redis client")
	}
	return &NewsfeedService{
		db:    db,
		redis: rd,
	}, nil
}
