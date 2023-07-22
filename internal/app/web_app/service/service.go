package service

import (
	"encoding/json"

	"github.com/EngineerProOrg/BE-K01/configs"
	authen_and_post2 "github.com/EngineerProOrg/BE-K01/pkg/client/authen_and_post"
	newsfeed2 "github.com/EngineerProOrg/BE-K01/pkg/client/newsfeed"
	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/authen_and_post"
	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/newsfeed"
	"go.uber.org/zap"
)

type WebService struct {
	authenticateAndPostClient authen_and_post.AuthenticateAndPostClient
	newsfeedClient            newsfeed.NewsfeedClient

	log *zap.Logger
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

	log, err := newLogger()
	if err != nil {
		return nil, err
	}
	return &WebService{
		authenticateAndPostClient: aapClient,
		newsfeedClient:            newsfeedClient,
		log:                       log,
	}, nil
}

func newLogger() (*zap.Logger, error) {
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "/tmp/logs"],
	  "errorOutputPaths": ["stderr"],
	  "initialFields": {"foo": "bar"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		return nil, err
	}
	logger := zap.Must(cfg.Build())
	return logger, nil
}
