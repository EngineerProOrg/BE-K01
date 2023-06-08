package handler

import (
	"github.com/EngineerProOrg/BE-K01/internal/user/model"
	"github.com/gin-gonic/gin"
	"crypto/rand"
    "encoding/base64"
	"time"
	
)


// 1 API /login, để tạo session cho mỗi người đăng nhập, dùng redis để lưu session id, user name ấy
func (hdl *userHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		sessionId, error := generateSessionID()
		if error != nil {
			ctx.JSON(400, gin.H{"error": error.Error()})
			return
		}
		if err := hdl.redis.Set(ctx, sessionId, user.Username, time.Duration(10000000000000)).Err() ; err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"sessionId": sessionId})
	}
}

func generateSessionID() (string, error) {
    // Generate 32 bytes of random data
    randomBytes := make([]byte, 32)
    _, err := rand.Read(randomBytes)

    if err != nil {
        return "", err
    }

    // Encode the random data to base64
    sessionID := base64.URLEncoding.EncodeToString(randomBytes)

    return sessionID, nil
}