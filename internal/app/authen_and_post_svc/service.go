package authen_and_post_svc

import (
	"context"

	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/authen_and_post"
)

func (a *AuthenticateAndPostService) CheckUserAuthentication(ctx context.Context, info *authen_and_post.UserInfo) (*authen_and_post.UserResult, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthenticateAndPostService) CreateUser(ctx context.Context, info *authen_and_post.UserDetailInfo) (*authen_and_post.UserResult, error) {
	//TODO implement me
	panic("implement me")
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
}

func NewAuthenticateAndPostService() *AuthenticateAndPostService {
	return &AuthenticateAndPostService{}
}
