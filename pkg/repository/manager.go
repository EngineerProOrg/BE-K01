package repository

import (
	"github.com/EngineerProOrg/BE-K01/pkg/types"
	"gorm.io/gorm"
)

type StudentRepository interface {
	GetStudentByIdx(id int64)
}

type repo struct {
	db *gorm.DB
}

func (r *repo) GetStudentByIdx(id int64) {
	std := &types.Student{}
	r.db.First(std, id)
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &repo{
		db: db,
	}
}
