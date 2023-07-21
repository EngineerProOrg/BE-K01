package service

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/EngineerProOrg/BE-K01/internal/pkg/types"
	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/authen_and_post"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (svc *WebService) GetFollowList(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if err != nil {
		svc.log.Error("invalid user id", zap.String("user_id", ctx.Param("user_id")))
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: "invalid user id"})
		return
	}
	resp, err := svc.authenticateAndPostClient.GetFollowerList(ctx, &authen_and_post.GetFollowerListRequest{
		UserId: userId,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.MessageResponse{Message: err.Error()})
		return
	}
	if resp.Status == authen_and_post.GetFollowerListResponse_NOT_FOUND {
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (svc *WebService) FollowUser(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if err != nil {
		svc.log.Error("invalid user id", zap.String("user_id", ctx.Param("user_id")))
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: "invalid user id"})
		return
	}
	sessionId, err := ctx.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		svc.log.Error("unauthorized")
		ctx.JSON(http.StatusUnauthorized, types.MessageResponse{Message: "unauthorized"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.MessageResponse{Message: "unexpected error"})
		return
	}
	currentUserId, err := strconv.ParseInt(sessionId, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.MessageResponse{Message: "unexpected error"})
		return
	}

	resp, err := svc.authenticateAndPostClient.FollowUser(ctx, &authen_and_post.FollowUserRequest{
		UserId:      currentUserId,
		FollowingId: userId,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.MessageResponse{Message: err.Error()})
		return
	}
	if resp.Status == authen_and_post.FollowUserResponse_NOT_FOUND {
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: "user not found"})
		return
	} else if resp.Status == authen_and_post.FollowUserResponse_ALREADY_FOLLOWED {
		ctx.JSON(http.StatusOK, types.MessageResponse{Message: "user already followed"})
		return
	}
	ctx.JSON(http.StatusOK, types.MessageResponse{Message: "follow user successfully"})
}

func (svc *WebService) UnfollowUser(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if err != nil {
		svc.log.Error("invalid user id", zap.String("user_id", ctx.Param("user_id")))
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: "invalid user id"})
		return
	}
	sessionId, err := ctx.Cookie("session_id")
	if errors.Is(err, http.ErrNoCookie) {
		svc.log.Error("unauthorized")
		ctx.JSON(http.StatusUnauthorized, types.MessageResponse{Message: "unauthorized"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.MessageResponse{Message: "unexpected error"})
		return
	}
	currentUserId, err := strconv.ParseInt(sessionId, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.MessageResponse{Message: "unexpected error"})
		return
	}

	resp, err := svc.authenticateAndPostClient.UnfollowUser(ctx, &authen_and_post.UnfollowUserRequest{
		UserId:      currentUserId,
		FollowingId: userId,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.MessageResponse{Message: err.Error()})
		return
	}
	if resp.Status == authen_and_post.UnfollowUserResponse_NOT_FOUND {
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: "user not found"})
		return
	} else if resp.Status == authen_and_post.UnfollowUserResponse_NOT_FOLLOWED {
		ctx.JSON(http.StatusOK, types.MessageResponse{Message: "user not followed"})
		return
	}
	ctx.JSON(http.StatusOK, types.MessageResponse{Message: "unfollow user successfully"})
}
