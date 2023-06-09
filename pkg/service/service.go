package service

import (
	"fmt"

	"github.com/EngineerProOrg/BE-K01/pkg/repository"
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
	db    *gorm.DB
}

type StudentManager struct {
	// Các trường dữ liệu khác của StudentManager (nếu có)
}

func (s *StudentManager) GetProfessor(id int64) interface{} {
	fmt.Println("vo day")
	return nil
}

func NewService(db *gorm.DB, redis *redis.Client) StudentManagerService {
	return &studentManagerService{
		repo:  repository.NewStudentRepository(db),
		redis: redis,
	}
}
