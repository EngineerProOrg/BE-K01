package v1

import (
	"github.com/EngineerProOrg/BE-K01/internal/app/web_app/service"
	"github.com/gin-gonic/gin"
)

func AddFriendRouter(r *gin.RouterGroup, svc *service.WebService) {
	friendRouter := r.Group("friends")

	friendRouter.GET(":user_id", svc.GetFollowList)
	friendRouter.POST(":user_id", svc.FollowUser)
	friendRouter.DELETE(":user_id", svc.UnfollowUser)
	friendRouter.GET(":user_id/posts", svc.GetUserPosts)
}
