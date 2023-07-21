package authen_and_post_svc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/EngineerProOrg/BE-K01/configs"
	"github.com/EngineerProOrg/BE-K01/internal/pkg/types"
	"github.com/EngineerProOrg/BE-K01/pkg/types/proto/pb/authen_and_post"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (a *AuthenticateAndPostService) CheckUserAuthentication(ctx context.Context, info *authen_and_post.CheckUserAuthenticationRequest) (*authen_and_post.CheckUserAuthenticationResponse, error) {
	var user types.User
	result := a.db.Where(&types.User{UserName: info.GetUserName()}).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &authen_and_post.CheckUserAuthenticationResponse{
			Status: authen_and_post.CheckUserAuthenticationResponse_NOT_FOUND,
		}, nil
	} else if result.Error != nil {
		return nil, result.Error
	}

	passwordWithSalt := []byte(info.UserPassword + user.Salt)
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), passwordWithSalt)
	if err != nil {
		return &authen_and_post.CheckUserAuthenticationResponse{
			Status: authen_and_post.CheckUserAuthenticationResponse_WRONG_PASSWORD,
		}, nil
	}

	return &authen_and_post.CheckUserAuthenticationResponse{
		Status: authen_and_post.CheckUserAuthenticationResponse_OK,
	}, nil
}

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateAlphabetSalt(length int) []byte {
	rand.Seed(time.Now().UnixNano())

	salt := make([]byte, length)
	for i := 0; i < length; i++ {
		salt[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return salt
}

func hashPassword(password string, salt []byte) (string, error) {
	// Append the salt to the password
	passwordWithSalt := []byte(password + string(salt))

	// Generate the bcrypt hash
	hash, err := bcrypt.GenerateFromPassword(passwordWithSalt, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (a *AuthenticateAndPostService) CreateUser(ctx context.Context, info *authen_and_post.UserDetailInfo) (*authen_and_post.UserResult, error) {
	salt := generateAlphabetSalt(16)
	hashedPassword, err := hashPassword(info.GetUserPassword(), salt)
	if err != nil {
		return nil, err
	}

	newUser := types.User{
		HashedPassword: hashedPassword,
		Salt:           string(salt),
		FirstName:      info.GetFirstName(),
		LastName:       info.GetLastName(),
		DateOfBirth:    info.Dob.AsTime(),
		Email:          info.GetEmail(),
		UserName:       info.GetUserName(),
	}
	a.db.Create(&newUser)
	return &authen_and_post.UserResult{
		Status: authen_and_post.UserStatus_OK,
		Info: &authen_and_post.UserDetailInfo{
			UserId:   int64(newUser.ID),
			UserName: newUser.UserName,
		},
	}, nil
}

// EditUser edit user info by looking up user id in mysql database and update it
func (a *AuthenticateAndPostService) EditUser(ctx context.Context, info *authen_and_post.EditUserRequest) (*authen_and_post.EditUserResponse, error) {
	var user types.User
	a.db.Where(&types.User{ID: uint(info.UserId)}).First(&user)
	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	if info.FirstName != nil {
		user.FirstName = info.GetFirstName()
	}
	if info.LastName != nil {
		user.LastName = info.GetLastName()
	}
	if info.UserPassword != nil {
		salt := generateAlphabetSalt(16)
		hashedPassword, err := hashPassword(info.GetUserPassword(), salt)
		if err != nil {
			return nil, err
		}
		user.HashedPassword = hashedPassword
		user.Salt = string(salt)
	}
	if info.Dob != nil {
		user.DateOfBirth = info.Dob.AsTime()
	}
	a.db.Save(&user)

	return &authen_and_post.EditUserResponse{
		UserId: int64(user.ID),
	}, nil
}

func (a *AuthenticateAndPostService) GetUserFollower(ctx context.Context, info *authen_and_post.UserInfo) (*authen_and_post.UserFollower, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthenticateAndPostService) GetPostDetail(ctx context.Context, request *authen_and_post.GetPostRequest) (*authen_and_post.Post, error) {
	//TODO implement me
	panic("implement me")
}

type AuthenticateAndPostService struct {
	authen_and_post.UnimplementedAuthenticateAndPostServer
	db    *gorm.DB
	redis *redis.Client

	log *zap.Logger
}

func NewAuthenticateAndPostService(conf *configs.AuthenticateAndPostConfig) (*AuthenticateAndPostService, error) {
	db, err := gorm.Open(mysql.New(conf.MySQL), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println("can not connect to db ", err)
		return nil, err
	}
	rd := redis.NewClient(&conf.Redis)
	if rd == nil {
		return nil, fmt.Errorf("can not init redis client")
	}

	log, err := newLogger()
	if err != nil {
		return nil, err
	}
	return &AuthenticateAndPostService{
		db:    db,
		redis: rd,
		log:   log,
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
