package authen_and_post

import (
	"context"
	"math/rand"

	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/authen_and_post"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type randomClient struct {
	clients []authen_and_post.AuthenticateAndPostClient
}

func (a *randomClient) CheckUserAuthentication(ctx context.Context, in *authen_and_post.CheckUserAuthenticationRequest, opts ...grpc.CallOption) (*authen_and_post.CheckUserAuthenticationResponse, error) {
	return a.clients[rand.Intn(len(a.clients))].CheckUserAuthentication(ctx, in, opts...)
}

func (a *randomClient) CreateUser(ctx context.Context, in *authen_and_post.UserDetailInfo, opts ...grpc.CallOption) (*authen_and_post.UserResult, error) {
	return a.clients[rand.Intn(len(a.clients))].CreateUser(ctx, in, opts...)
}

func (a *randomClient) EditUser(ctx context.Context, in *authen_and_post.EditUserRequest, opts ...grpc.CallOption) (*authen_and_post.EditUserResponse, error) {
	return a.clients[rand.Intn(len(a.clients))].EditUser(ctx, in, opts...)
}

func (a *randomClient) GetUserFollower(ctx context.Context, in *authen_and_post.UserInfo, opts ...grpc.CallOption) (*authen_and_post.UserFollower, error) {
	return a.clients[rand.Intn(len(a.clients))].GetUserFollower(ctx, in, opts...)
}

func (a *randomClient) GetPostDetail(ctx context.Context, in *authen_and_post.GetPostRequest, opts ...grpc.CallOption) (*authen_and_post.Post, error) {
	return a.clients[rand.Intn(len(a.clients))].GetPostDetail(ctx, in, opts...)
}

func (a *randomClient) FollowUser(ctx context.Context, in *authen_and_post.FollowUserRequest, opts ...grpc.CallOption) (*authen_and_post.FollowUserResponse, error) {
	return a.clients[rand.Intn(len(a.clients))].FollowUser(ctx, in, opts...)
}

func (a *randomClient) UnfollowUser(ctx context.Context, in *authen_and_post.UnfollowUserRequest, opts ...grpc.CallOption) (*authen_and_post.UnfollowUserResponse, error) {
	return a.clients[rand.Intn(len(a.clients))].UnfollowUser(ctx, in, opts...)
}

func (a *randomClient) GetFollowerList(ctx context.Context, in *authen_and_post.GetFollowerListRequest, opts ...grpc.CallOption) (*authen_and_post.GetFollowerListResponse, error) {
	return a.clients[rand.Intn(len(a.clients))].GetFollowerList(ctx, in, opts...)
}

func (a *randomClient) GetPost(ctx context.Context, in *authen_and_post.GetPostRequest, opts ...grpc.CallOption) (*authen_and_post.GetPostResponse, error) {
	return a.clients[rand.Intn(len(a.clients))].GetPost(ctx, in, opts...)
}

func (a *randomClient) CreatePost(ctx context.Context, in *authen_and_post.CreatePostRequest, opts ...grpc.CallOption) (*authen_and_post.CreatePostResponse, error) {
	return a.clients[rand.Intn(len(a.clients))].CreatePost(ctx, in, opts...)
}

func (a *randomClient) DeletePost(ctx context.Context, in *authen_and_post.DeletePostRequest, opts ...grpc.CallOption) (*authen_and_post.DeletePostResponse, error) {
	return a.clients[rand.Intn(len(a.clients))].DeletePost(ctx, in, opts...)
}

func (a *randomClient) EditPost(ctx context.Context, in *authen_and_post.EditPostRequest, opts ...grpc.CallOption) (*authen_and_post.EditPostResponse, error) {
	return a.clients[rand.Intn(len(a.clients))].EditPost(ctx, in, opts...)
}

func (a *randomClient) CreatePostComment(ctx context.Context, in *authen_and_post.CreatePostCommentRequest, opts ...grpc.CallOption) (*authen_and_post.CreatePostCommentResponse, error) {
	return a.clients[rand.Intn(len(a.clients))].CreatePostComment(ctx, in, opts...)
}

func (a *randomClient) LikePost(ctx context.Context, in *authen_and_post.LikePostRequest, opts ...grpc.CallOption) (*authen_and_post.LikePostResponse, error) {
	return a.clients[rand.Intn(len(a.clients))].LikePost(ctx, in, opts...)
}

func NewClient(hosts []string) (authen_and_post.AuthenticateAndPostClient, error) {
	var opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	clients := make([]authen_and_post.AuthenticateAndPostClient, 0, len(hosts))
	for _, host := range hosts {
		conn, err := grpc.Dial(host, opts...)
		if err != nil {
			return nil, err
		}
		client := authen_and_post.NewAuthenticateAndPostClient(conn)
		clients = append(clients, client)
	}
	return &randomClient{clients}, nil
}
