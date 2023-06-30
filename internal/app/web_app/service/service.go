package service

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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

// CheckUserNamePassword godoc
//
//	@Summary		get user
//	@Description	check user user_name and password
//	@Tags			test
//	@Accept			json
//	@Produce		json
//	@Param			request body types.LoginRequest true "login param"
//	@Success		200	{object} types.MessageResponse
//	@Failure		400	{object} types.MessageResponse
//	@Failure		500	{object} types.MessageResponse
//	@Router			/users/login [post]
func (svc *WebService) CheckUserNamePassword(ctx *gin.Context) {
	start := time.Now()
	status := http.StatusOK
	countExporter.WithLabelValues("check_user_login", "total").Inc()
	defer func() {
		latencyExporter.WithLabelValues("check_user_login", strconv.Itoa(status)).Observe(float64(start.UnixMilli()))
	}()
	var jsonRequest types.LoginRequest
	err := ctx.ShouldBindJSON(&jsonRequest)
	if err != nil {
		countExporter.WithLabelValues("check_user_login", "bad_request").Inc()
		status = http.StatusBadRequest
		ctx.JSON(status, &types.MessageResponse{Message: err.Error()})
		return
	}
	authentication, err := svc.authenticateAndPostClient.CheckUserAuthentication(ctx, &authen_and_post.UserInfo{
		UserName:     jsonRequest.UserName,
		UserPassword: jsonRequest.Password,
	})
	if err != nil {
		countExporter.WithLabelValues("check_user_login", "call_api_failed").Inc()
		status = http.StatusInternalServerError
		ctx.JSON(status, &types.MessageResponse{Message: err.Error()})
		return
	}
	if authentication.GetStatus() == authen_and_post.UserStatus_OK {
		countExporter.WithLabelValues("check_user_login", "success").Inc()
		ctx.JSON(status, &types.MessageResponse{Message: "ok"})
		// change this later
		ctx.SetCookie("session_id", fmt.Sprintf("%d", authentication.Info.UserId), 0, "", "", false, false)
	}
}

func (svc *WebService) CreateUser(ctx *gin.Context) {

}

func (svc *WebService) EditUser(ctx *gin.Context) {

}
