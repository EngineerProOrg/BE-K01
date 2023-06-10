package newsfeed

import (
	"context"
	"math/rand"

	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/newsfeed"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type randomClient struct {
	clients []newsfeed.NewsfeedClient
}

func (r *randomClient) Newsfeed(ctx context.Context, in *newsfeed.NewsfeedRequest, opts ...grpc.CallOption) (*newsfeed.NewsfeedResponse, error) {
	return r.clients[rand.Intn(len(r.clients))].Newsfeed(ctx, in, opts...)
}

func NewClient(hosts []string) (newsfeed.NewsfeedClient, error) {
	var opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	clients := make([]newsfeed.NewsfeedClient, 0, len(hosts))
	for _, host := range hosts {
		conn, err := grpc.Dial(host, opts...)
		if err != nil {
			return nil, err
		}
		client := newsfeed.NewNewsfeedClient(conn)
		clients = append(clients, client)
	}
	return &randomClient{clients}, nil
}
