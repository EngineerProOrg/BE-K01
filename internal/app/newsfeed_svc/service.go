package newsfeed_svc

import (
	"context"

	"github.com/EngineerProOrg/BE-K01/configs"
	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/newsfeed"
)

type NewsfeedService struct {
	newsfeed.UnimplementedNewsfeedServer
}

func (n NewsfeedService) Newsfeed(ctx context.Context, request *newsfeed.NewsfeedRequest) (*newsfeed.NewsfeedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewNewsfeedService(config *configs.NewsfeedConfig) *NewsfeedService {
	return &NewsfeedService{}
}
