package student_service

import (
	"github.com/EngineerProOrg/BE-K01/pkg/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type StudentManagerService interface {
	GetStudent(id int64) interface{}
	GetProfessor(id int64) interface{}
}

type studentManagerService struct {
	repo  repository.StudentRepository
	redis *redis.Client
}

func NewService(db *gorm.DB, redis *redis.Client) StudentManagerService {
	return &studentManagerService{
		repo:  repository.NewStudentRepository(db),
		redis: redis,
	}
}

func MappingService(r *gin.Engine, service StudentManagerService) {
	r.GET("test", func(context *gin.Context) {
		service.GetProfessor(1)
	})
}
