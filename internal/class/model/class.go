package classmodel

import (
	"github.com/EngineerProOrg/BE-K01/internal/room/model"
	"github.com/EngineerProOrg/BE-K01/internal/course/model"
	"github.com/EngineerProOrg/BE-K01/internal/professor/model"
)

type Class struct {
	Class_id int `json:"class_id" gorm:"column:CLASS_ID;primaryKey"`
	Class_name string `json:"class_name" gorm:"column:CLASS_NAME"`
	Prof_id int `json:"prof_id" gorm:"column:PROF_ID"`
	Prof professormodel.Professor `json:"prof" gorm:"preload:false"`
	Course_id int `json:"course_id" gorm:"column:COURSE_ID"`
	Course coursemodel.Course `json:"course" gorm:"preload:false"`
	Room_id int `json:"room_id" gorm:"column:ROOM_ID"`
	Room roommodel.Room `json:"room" gorm:"preload:false"`
}

func (Class) TableName() string {
	return "CLASS"
}