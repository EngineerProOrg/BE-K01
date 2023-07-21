package service

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EngineerProOrg/BE-K01/internal/pkg/types"
	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/authen_and_post"
	"github.com/gin-gonic/gin"
)

func (svc *WebService) GetPost(ctx *gin.Context) {
	postId, err := strconv.ParseInt(ctx.Param("post_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: fmt.Sprintf("invalid post id: %s", ctx.Param("post_id"))})
		return
	}
	resp, err := svc.authenticateAndPostClient.GetPost(ctx, &authen_and_post.GetPostRequest{
		PostId: postId,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.MessageResponse{Message: err.Error()})
		return
	}
	if resp.Status == authen_and_post.GetPostResponse_NOT_FOUND {
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: "post not found"})
		return
	}
	ctx.JSON(http.StatusOK, types.PostDetailResponse{
		PostID:           resp.Post.PostId,
		UserID:           resp.Post.UserId,
		ContentText:      resp.Post.ContentText,
		ContentImagePath: resp.Post.ContentImagePath,
		Visible:          resp.Post.Visible,
		CreatedTime:      resp.Post.CreatedTime.AsTime(),
	})
}

func (svc *WebService) CreatePost(ctx *gin.Context) {
	var jsonRequest types.CreatePostRequest
	if err := ctx.ShouldBindJSON(&jsonRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: err.Error()})
		return
	}
	resp, err := svc.authenticateAndPostClient.CreatePost(ctx, &authen_and_post.CreatePostRequest{
		UserId:           jsonRequest.UserId,
		ContentText:      jsonRequest.ContentText,
		ContentImagePath: jsonRequest.ContentImagePath,
		Visible:          jsonRequest.Visible,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.MessageResponse{Message: err.Error()})
		return
	}
	if resp.Status == authen_and_post.CreatePostResponse_NOT_FOUND {
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, types.MessageResponse{Message: fmt.Sprintf("create post successfully with id: %d", resp.PostId)})
}

func (svc *WebService) EditPost(ctx *gin.Context) {
	postId, err := strconv.ParseInt(ctx.Param("post_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: fmt.Sprintf("invalid post id: %s", ctx.Param("post_id"))})
		return
	}
	var jsonRequest types.EditPostRequest
	if err := ctx.ShouldBindJSON(&jsonRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: err.Error()})
		return
	}
	req := &authen_and_post.EditPostRequest{}
	req.PostId = postId
	if jsonRequest.ContentText != nil {
		req.ContentText = jsonRequest.ContentText
	}
	if jsonRequest.ContentImagePath != nil {
		req.ContentImagePath = jsonRequest.ContentImagePath
	}
	if jsonRequest.Visible != nil {
		req.Visible = jsonRequest.Visible
	}
	resp, err := svc.authenticateAndPostClient.EditPost(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.MessageResponse{Message: err.Error()})
		return
	}
	if resp.Status == authen_and_post.EditPostResponse_NOT_FOUND {
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: "post not found"})
		return
	}
	ctx.JSON(http.StatusOK, types.MessageResponse{Message: fmt.Sprintf("edit post successfully with id: %d", postId)})
}

func (svc *WebService) DeletePost(ctx *gin.Context) {
	postId, err := strconv.ParseInt(ctx.Param("post_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: fmt.Sprintf("invalid post id: %s", ctx.Param("post_id"))})
		return
	}
	resp, err := svc.authenticateAndPostClient.DeletePost(ctx, &authen_and_post.DeletePostRequest{
		PostId: postId,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.MessageResponse{Message: err.Error()})
		return
	}
	if resp.Status == authen_and_post.DeletePostResponse_NOT_FOUND {
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: "post not found"})
		return
	}
	ctx.JSON(http.StatusOK, types.MessageResponse{Message: fmt.Sprintf("delete post successfully with id: %d", postId)})
}

func (svc *WebService) CreatePostComment(ctx *gin.Context) {
	postId, err := strconv.ParseInt(ctx.Param("post_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: fmt.Sprintf("invalid post id: %s", ctx.Param("post_id"))})
		return
	}
	var jsonRequest types.CreatePostCommentRequest
	if err := ctx.ShouldBindJSON(&jsonRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: err.Error()})
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
	resp, err := svc.authenticateAndPostClient.CreatePostComment(ctx, &authen_and_post.CreatePostCommentRequest{
		PostId:      postId,
		UserId:      currentUserId,
		ContentText: jsonRequest.ContentText,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.MessageResponse{Message: err.Error()})
		return
	}
	if resp.Status == authen_and_post.CreatePostCommentResponse_POST_NOT_FOUND {
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: "post not found"})
		return
	} else if resp.Status == authen_and_post.CreatePostCommentResponse_USER_NOT_FOUND {
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, types.MessageResponse{Message: fmt.Sprintf("create post comment successfully with id: %d", resp.CommentId)})
}

func (svc *WebService) LikePost(ctx *gin.Context) {
	postId, err := strconv.ParseInt(ctx.Param("post_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: fmt.Sprintf("invalid post id: %s", ctx.Param("post_id"))})
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
	resp, err := svc.authenticateAndPostClient.LikePost(ctx, &authen_and_post.LikePostRequest{
		PostId: postId,
		UserId: currentUserId,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.MessageResponse{Message: err.Error()})
		return
	}
	if resp.Status == authen_and_post.LikePostResponse_POST_NOT_FOUND {
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: "post not found"})
		return
	} else if resp.Status == authen_and_post.LikePostResponse_USER_NOT_FOUND {
		ctx.JSON(http.StatusBadRequest, types.MessageResponse{Message: "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, types.MessageResponse{Message: fmt.Sprintf("like post successfully with id: %d", postId)})
}
