package service

import (
	"github.com/EngineerProOrg/BE-K01/configs"
	authen_and_post2 "github.com/EngineerProOrg/BE-K01/pkg/client/authen_and_post"
	newsfeed2 "github.com/EngineerProOrg/BE-K01/pkg/client/newsfeed"
	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/authen_and_post"
	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/newsfeed"
)

type WebService struct {
	authenticateAndPostClient authen_and_post.AuthenticateAndPostClient
	newsfeedClient            newsfeed.NewsfeedClient
}

func NewWebService(conf *configs.WebConfig) (*WebService, error) {
	aapClient, err := authen_and_post2.NewClient(conf.AuthenticateAndPost.Hosts)
	if err != nil {
		return nil, err
	}
	newsfeedClient, err := newsfeed2.NewClient(conf.Newsfeed.Hosts)
	if err != nil {
		return nil, err
	}
	return &WebService{
		authenticateAndPostClient: aapClient,
		newsfeedClient:            newsfeedClient,
	}, nil
}
