package types

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	FirstName string `gorm:"first_name"`
	LastName  string `gorm:"last_name"`
}

func (Student) TableName() string {
	return "student"
}
