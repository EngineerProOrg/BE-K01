package authen_and_post_svc

import (
	"context"
	"errors"

	"github.com/EngineerProOrg/BE-K01/internal/pkg/types"
	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/authen_and_post"
	"go.uber.org/zap"
)

func (a *AuthenticateAndPostService) ensureUserExist(ctx context.Context, userId int64) error {
	var user types.User
	err := a.db.Table("user").Where("id = ?", userId).First(&user).Error
	if err != nil {
		return errors.New("user not found")
	}
	return nil
}

func (a *AuthenticateAndPostService) FollowUser(ctx context.Context, req *authen_and_post.FollowUserRequest) (*authen_and_post.FollowUserResponse, error) {
	err := a.ensureUserExist(ctx, req.UserId)
	if err != nil {
		return &authen_and_post.FollowUserResponse{Status: authen_and_post.FollowUserResponse_NOT_FOUND}, nil
	}
	err = a.ensureUserExist(ctx, req.FollowingId)
	if err != nil {
		return &authen_and_post.FollowUserResponse{Status: authen_and_post.FollowUserResponse_NOT_FOUND}, nil
	}

	var user types.User
	err = a.db.Preload("Following").Preload("Follower").Find(&user, req.UserId).Error
	a.log.Debug("returned user", zap.Any("user", user))
	var alreadyFollowed bool
	for _, following := range user.Following {
		if following.ID == uint(req.FollowingId) {
			alreadyFollowed = true
			break
		}
	}
	if !alreadyFollowed {
		a.log.Info("following new user")
		var followingUser types.User
		a.db.Where(&types.User{ID: uint(req.FollowingId)}).First(&followingUser)

		err := a.db.Model(&user).Association("Following").Append(&followingUser)
		if err != nil {
			return nil, err
		}
		err = a.db.Model(&followingUser).Association("Follower").Append(&user)
		if err != nil {
			return nil, err
		}
		return &authen_and_post.FollowUserResponse{Status: authen_and_post.FollowUserResponse_OK}, nil
	} else {
		return &authen_and_post.FollowUserResponse{Status: authen_and_post.FollowUserResponse_ALREADY_FOLLOWED}, nil
	}
}

func (a *AuthenticateAndPostService) UnfollowUser(ctx context.Context, req *authen_and_post.UnfollowUserRequest) (*authen_and_post.UnfollowUserResponse, error) {
	err := a.ensureUserExist(ctx, req.UserId)
	if err != nil {
		return &authen_and_post.UnfollowUserResponse{Status: authen_and_post.UnfollowUserResponse_NOT_FOUND}, nil
	}
	err = a.ensureUserExist(ctx, req.FollowingId)
	if err != nil {
		return &authen_and_post.UnfollowUserResponse{Status: authen_and_post.UnfollowUserResponse_NOT_FOUND}, nil
	}

	var user types.User
	err = a.db.Preload("Following").Preload("Follower").Find(&user, req.UserId).Error
	a.log.Debug("returned user", zap.Any("user", user))
	var alreadyFollowed bool
	for _, following := range user.Following {
		if following.ID == uint(req.FollowingId) {
			alreadyFollowed = true
			break
		}
	}
	if alreadyFollowed {
		a.log.Info("unfollowing user")
		var followingUser types.User
		a.db.Where(&types.User{ID: uint(req.FollowingId)}).First(&followingUser)

		err := a.db.Model(&user).Association("Following").Delete(&followingUser)
		if err != nil {
			return nil, err
		}
		err = a.db.Model(&followingUser).Association("Follower").Delete(&user)
		if err != nil {
			return nil, err
		}
		return &authen_and_post.UnfollowUserResponse{Status: authen_and_post.UnfollowUserResponse_OK}, nil
	} else {
		return &authen_and_post.UnfollowUserResponse{Status: authen_and_post.UnfollowUserResponse_NOT_FOLLOWED}, nil
	}
}

func (a *AuthenticateAndPostService) GetFollowerList(ctx context.Context, req *authen_and_post.GetFollowerListRequest) (*authen_and_post.GetFollowerListResponse, error) {
	err := a.ensureUserExist(ctx, req.UserId)
	if err != nil {
		return &authen_and_post.GetFollowerListResponse{Status: authen_and_post.GetFollowerListResponse_NOT_FOUND}, nil
	}

	var user types.User
	err = a.db.Preload("Follower").Find(&user, req.UserId).Error
	a.log.Debug("returned user", zap.Any("user", user))
	var followersInfo []*authen_and_post.GetFollowerListResponse_FollowerInfo
	for _, follower := range user.Follower {
		followersInfo = append(followersInfo, &authen_and_post.GetFollowerListResponse_FollowerInfo{
			UserId:   int64(follower.ID),
			UserName: follower.UserName,
		})
	}
	return &authen_and_post.GetFollowerListResponse{Status: authen_and_post.GetFollowerListResponse_OK, Followers: followersInfo}, nil
}
