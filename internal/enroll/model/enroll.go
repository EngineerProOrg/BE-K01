package enrollmodel

import (
	"github.com/EngineerProOrg/BE-K01/internal/class/model"
	"github.com/EngineerProOrg/BE-K01/internal/student/model"
)

// create table ENROLL(
//     STUD_ID int,
//     CLASS_ID int,
//     GRADE varchar(3),
//     foreign key (STUD_ID) references STUDENT(STUD_ID),
//     foreign key (CLASS_ID) references CLASS(CLASS_ID),
//     PRIMARY KEY (STUD_ID, CLASS_ID)
// )

type Enroll struct {
	Stud_id int `json:"stud_id" gorm:"column:STUD_ID;primaryKey"`
	Stud studentmodel.Student `json:"stud" gorm:"preload:false"`
	Class_id int `json:"class_id" gorm:"column:CLASS_ID;primaryKey"`
	Class classmodel.Class `json:"class" gorm:"preload:false"`
	Grade string `json:"grade" gorm:"column:GRADE"`
}

func (Enroll) TableName() string {
	return "ENROLL"
}