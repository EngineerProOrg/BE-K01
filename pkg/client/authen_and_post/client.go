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

func (a *randomClient) CheckUserAuthentication(ctx context.Context, in *authen_and_post.UserInfo, opts ...grpc.CallOption) (*authen_and_post.UserResult, error) {
	return a.clients[rand.Intn(len(a.clients))].CheckUserAuthentication(ctx, in, opts...)
}

func (a *randomClient) CreateUser(ctx context.Context, in *authen_and_post.UserDetailInfo, opts ...grpc.CallOption) (*authen_and_post.UserResult, error) {
	return a.clients[rand.Intn(len(a.clients))].CreateUser(ctx, in, opts...)
}

func (a *randomClient) EditUser(ctx context.Context, in *authen_and_post.UserDetailInfo, opts ...grpc.CallOption) (*authen_and_post.UserResult, error) {
	return a.clients[rand.Intn(len(a.clients))].EditUser(ctx, in, opts...)
}

func (a *randomClient) GetUserFollower(ctx context.Context, in *authen_and_post.UserInfo, opts ...grpc.CallOption) (*authen_and_post.UserFollower, error) {
	return a.clients[rand.Intn(len(a.clients))].GetUserFollower(ctx, in, opts...)
}

func (a *randomClient) GetPostDetail(ctx context.Context, in *authen_and_post.GetPostRequest, opts ...grpc.CallOption) (*authen_and_post.Post, error) {
	return a.clients[rand.Intn(len(a.clients))].GetPostDetail(ctx, in, opts...)
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
