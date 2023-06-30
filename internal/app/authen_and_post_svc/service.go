package authen_and_post_svc

import (
	"context"
	"fmt"

	"github.com/EngineerProOrg/BE-K01/configs"
	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/authen_and_post"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (a *AuthenticateAndPostService) CheckUserAuthentication(ctx context.Context, info *authen_and_post.UserInfo) (*authen_and_post.UserResult, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthenticateAndPostService) CreateUser(ctx context.Context, info *authen_and_post.UserDetailInfo) (*authen_and_post.UserResult, error) {
	return &authen_and_post.UserResult{
		Status: authen_and_post.UserStatus_OK,
		Info: &authen_and_post.UserDetailInfo{
			UserId:   1,
			UserName: "dong.truong",
		},
	}, nil
}

func (a *AuthenticateAndPostService) EditUser(ctx context.Context, info *authen_and_post.UserDetailInfo) (*authen_and_post.UserResult, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthenticateAndPostService) GetUserFollower(ctx context.Context, info *authen_and_post.UserInfo) (*authen_and_post.UserFollower, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthenticateAndPostService) GetPostDetail(ctx context.Context, request *authen_and_post.GetPostRequest) (*authen_and_post.Post, error) {
	//TODO implement me
	panic("implement me")
}

type AuthenticateAndPostService struct {
	authen_and_post.UnimplementedAuthenticateAndPostServer
	db    *gorm.DB
	redis *redis.Client
}

func NewAuthenticateAndPostService(conf *configs.AuthenticateAndPostConfig) (*AuthenticateAndPostService, error) {
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
	return &AuthenticateAndPostService{
		db:    db,
		redis: rd,
	}, nil
}
