package service

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/EngineerProOrg/BE-K01/internal/pkg/types"
	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/authen_and_post"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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
	authentication, err := svc.authenticateAndPostClient.CheckUserAuthentication(ctx, &authen_and_post.CheckUserAuthenticationRequest{
		UserName:     jsonRequest.UserName,
		UserPassword: jsonRequest.Password,
	})
	if err != nil {
		countExporter.WithLabelValues("check_user_login", "call_api_failed").Inc()
		status = http.StatusInternalServerError
		ctx.JSON(status, &types.MessageResponse{Message: err.Error()})
		return
	}
	if authentication.GetStatus() == authen_and_post.CheckUserAuthenticationResponse_OK {
		countExporter.WithLabelValues("check_user_login", "success").Inc()
		ctx.JSON(status, &types.MessageResponse{Message: "ok"})
		// change this later
		ctx.SetCookie("session_id", fmt.Sprintf("%d", authentication.UserId), 0, "", "", false, false)
	} else if authentication.GetStatus() == authen_and_post.CheckUserAuthenticationResponse_NOT_FOUND {
		countExporter.WithLabelValues("check_user_login", "not_found").Inc()
		ctx.JSON(status, &types.MessageResponse{Message: "not found"})
	} else {
		countExporter.WithLabelValues("check_user_login", "wrong_password").Inc()
		ctx.JSON(status, &types.MessageResponse{Message: "wrong password"})
	}
}

func (svc *WebService) CreateUser(ctx *gin.Context) {
	var jsonRequest types.CreateUserRequest
	err := ctx.ShouldBindJSON(&jsonRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &types.MessageResponse{Message: err.Error()})
		return
	}
	dob, err := time.Parse("2006-01-02", jsonRequest.Dob)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &types.MessageResponse{Message: err.Error()})
		return
	}
	createdUser, err := svc.authenticateAndPostClient.CreateUser(ctx, &authen_and_post.UserDetailInfo{
		FirstName:    jsonRequest.FirstName,
		LastName:     jsonRequest.LastName,
		Dob:          timestamppb.New(dob),
		UserName:     jsonRequest.UserName,
		UserPassword: jsonRequest.Password,
		Email:        jsonRequest.Email,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &types.MessageResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &types.MessageResponse{Message: fmt.Sprintf("Successfully created user with id: %d", createdUser.Info.UserId)})
}

func (svc *WebService) EditUser(ctx *gin.Context) {
	var jsonRequest types.EditUserRequest
	err := ctx.ShouldBindJSON(&jsonRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &types.MessageResponse{Message: err.Error()})
		return
	}
	userId := jsonRequest.UserId
	if userId == 0 {
		ctx.JSON(http.StatusBadRequest, &types.MessageResponse{Message: "User id is required"})
		return
	}
	var firstName *string
	if jsonRequest.FirstName != "" {
		firstName = &jsonRequest.FirstName
	}
	var lastName *string
	if jsonRequest.LastName != "" {
		lastName = &jsonRequest.LastName
	}
	var dob *timestamppb.Timestamp
	if jsonRequest.Dob != "" {
		parsedDob, err := time.Parse("2006-01-02", jsonRequest.Dob)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, &types.MessageResponse{Message: err.Error()})
			return
		}
		dob = timestamppb.New(parsedDob)
	}
	var password *string
	if jsonRequest.Password != "" {
		password = &jsonRequest.Password
	}

	resp, err := svc.authenticateAndPostClient.EditUser(ctx, &authen_and_post.EditUserRequest{
		UserId:       userId,
		FirstName:    firstName,
		LastName:     lastName,
		Dob:          dob,
		UserPassword: password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &types.MessageResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &types.MessageResponse{Message: "Successfully edited user with id: " + fmt.Sprintf("%d", resp.UserId)})
}
