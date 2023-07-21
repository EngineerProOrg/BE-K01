package authen_and_post_svc

import (
	"context"
	"errors"

	"github.com/EngineerProOrg/BE-K01/internal/pkg/types"
	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/authen_and_post"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func (a *AuthenticateAndPostService) GetPost(ctx context.Context, in *authen_and_post.GetPostRequest) (*authen_and_post.GetPostResponse, error) {
	a.log.Debug("start get post")
	defer a.log.Debug("end get post")

	var post types.Post
	err := a.db.First(&post, in.PostId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &authen_and_post.GetPostResponse{
			Status: authen_and_post.GetPostResponse_NOT_FOUND,
		}, nil
	}
	if err != nil {
		return nil, err
	}
	return &authen_and_post.GetPostResponse{
		Status: authen_and_post.GetPostResponse_OK,
		Post: &authen_and_post.Post{
			PostId:           int64(post.ID),
			UserId:           int64(post.UserID),
			ContentText:      post.ContentText,
			ContentImagePath: post.ContentImagePath,
			Visible:          post.Visible,
			CreatedTime:      timestamppb.New(post.CreatedAt),
		},
	}, nil
}

func (a *AuthenticateAndPostService) CreatePost(ctx context.Context, in *authen_and_post.CreatePostRequest) (*authen_and_post.CreatePostResponse, error) {
	a.log.Debug("start create post")
	defer a.log.Debug("end create post")

	err := a.ensureUserExist(ctx, in.UserId)
	if err != nil {
		return &authen_and_post.CreatePostResponse{
			Status: authen_and_post.CreatePostResponse_NOT_FOUND,
		}, nil
	}

	var user types.User
	err = a.db.Preload("Posts").Find(&user, in.UserId).Error
	if err != nil {
		return nil, err
	}

	post := types.Post{
		ContentText:      in.ContentText,
		ContentImagePath: in.ContentImagePath,
		UserID:           uint(in.UserId),
		Visible:          in.Visible,
	}

	err = a.db.Model(&user).Association("Posts").Append(&post)
	if err != nil {
		return nil, err
	}
	return &authen_and_post.CreatePostResponse{Status: authen_and_post.CreatePostResponse_OK, PostId: int64(post.ID)}, nil
}

func (a *AuthenticateAndPostService) DeletePost(ctx context.Context, in *authen_and_post.DeletePostRequest) (*authen_and_post.DeletePostResponse, error) {
	a.log.Debug("start delete post")
	defer a.log.Debug("end delete post")

	var post types.Post
	err := a.db.First(&post, in.PostId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &authen_and_post.DeletePostResponse{
			Status: authen_and_post.DeletePostResponse_NOT_FOUND,
		}, nil
	}
	if err != nil {
		return nil, err
	}
	err = a.db.Delete(&post).Error
	if err != nil {
		return nil, err
	}
	return &authen_and_post.DeletePostResponse{Status: authen_and_post.DeletePostResponse_OK}, nil
}

// EditPost edit post
func (a *AuthenticateAndPostService) EditPost(ctx context.Context, in *authen_and_post.EditPostRequest) (*authen_and_post.EditPostResponse, error) {
	a.log.Debug("start edit post", zap.Any("in", in))
	defer a.log.Debug("end edit post")

	var post types.Post
	err := a.db.First(&post, in.PostId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &authen_and_post.EditPostResponse{
			Status: authen_and_post.EditPostResponse_NOT_FOUND,
		}, nil
	}
	if err != nil {
		return nil, err
	}
	a.log.Debug("post", zap.Any("post", post))

	if in.ContentText != nil {
		post.ContentText = in.GetContentText()
	}
	if in.ContentImagePath != nil {
		post.ContentImagePath = in.GetContentImagePath()
	}
	if in.Visible != nil {
		post.Visible = in.GetVisible()
	}
	a.log.Debug("updated post", zap.Any("post", post))

	err = a.db.Save(&post).Error
	if err != nil {
		return nil, err
	}
	return &authen_and_post.EditPostResponse{Status: authen_and_post.EditPostResponse_OK}, nil
}

// CreatePostComment create post comment
func (a *AuthenticateAndPostService) CreatePostComment(ctx context.Context, in *authen_and_post.CreatePostCommentRequest) (*authen_and_post.CreatePostCommentResponse, error) {
	a.log.Debug("start create post comment")
	defer a.log.Debug("end create post comment")

	err := a.ensureUserExist(ctx, in.UserId)
	if err != nil {
		return &authen_and_post.CreatePostCommentResponse{
			Status: authen_and_post.CreatePostCommentResponse_USER_NOT_FOUND,
		}, nil
	}

	var post types.Post
	err = a.db.Where("id = ?", in.PostId).First(&post).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		a.log.Debug("post not found")
		return &authen_and_post.CreatePostCommentResponse{
			Status: authen_and_post.CreatePostCommentResponse_POST_NOT_FOUND,
		}, nil
	}
	if err != nil {
		return nil, err
	}
	a.log.Debug("post found", zap.Any("post", post))
	a.db.Preload("Comments").Find(&post, in.PostId)

	comment := types.Comment{
		Content: in.ContentText,
		UserID:  uint(in.UserId),
		PostID:  uint(in.PostId),
	}

	err = a.db.Model(&post).Association("Comments").Append(&comment)
	if err != nil {
		return nil, err
	}
	return &authen_and_post.CreatePostCommentResponse{Status: authen_and_post.CreatePostCommentResponse_OK, CommentId: int64(comment.ID)}, nil
}

// LikePost like post
func (a *AuthenticateAndPostService) LikePost(ctx context.Context, in *authen_and_post.LikePostRequest) (*authen_and_post.LikePostResponse, error) {
	a.log.Debug("start like post")
	defer a.log.Debug("end like post")

	var err error
	var user types.User
	err = a.db.First(&user, in.UserId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &authen_and_post.LikePostResponse{
			Status: authen_and_post.LikePostResponse_USER_NOT_FOUND,
		}, nil
	}

	var post types.Post
	err = a.db.First(&post, in.PostId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &authen_and_post.LikePostResponse{
			Status: authen_and_post.LikePostResponse_POST_NOT_FOUND,
		}, nil
	}
	if err != nil {
		return nil, err
	}

	err = a.db.Preload("LikedUsers").Find(&post, in.PostId).Error
	if err != nil {
		return nil, err
	}

	err = a.db.Model(&post).Association("LikedUsers").Append(&user)
	if err != nil {
		return nil, err
	}
	return &authen_and_post.LikePostResponse{Status: authen_and_post.LikePostResponse_OK}, nil
}
