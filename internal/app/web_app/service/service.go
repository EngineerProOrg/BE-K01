package service

import (
	"fmt"
	"net/http"

	"github.com/EngineerProOrg/BE-K01/configs"
	"github.com/EngineerProOrg/BE-K01/internal/pkg/types"
	authen_and_post2 "github.com/EngineerProOrg/BE-K01/pkg/client/authen_and_post"
	newsfeed2 "github.com/EngineerProOrg/BE-K01/pkg/client/newsfeed"
	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/authen_and_post"
	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/newsfeed"
	"github.com/gin-gonic/gin"
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

func (svc *WebService) CheckUserNamePassword(ctx *gin.Context) {
	var jsonRequest types.LoginRequest
	err := ctx.ShouldBindJSON(jsonRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	authentication, err := svc.authenticateAndPostClient.CheckUserAuthentication(ctx, &authen_and_post.UserInfo{
		UserName:     jsonRequest.UserName,
		UserPassword: jsonRequest.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if authentication.GetStatus() == authen_and_post.UserStatus_OK {
		ctx.Status(http.StatusOK)
		// change this later
		ctx.SetCookie("session_id", fmt.Sprintf("%d", authentication.Info.UserId), 0, "", "", false, false)
	}
}

func (svc *WebService) CreateUser(ctx *gin.Context) {

}

func (svc *WebService) EditUser(ctx *gin.Context) {

}
